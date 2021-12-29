import logging
import threading
import datetime
import copy
import configparser

import pandas as pd
from sqlalchemy import create_engine, Column, Integer, String, ForeignKey, desc
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.exc import IntegrityError
from contextlib import contextmanager


config = configparser.ConfigParser()
config.read('setting.ini')
default = config['DEFAULT']
logger = logging.getLogger(__name__)
lock = threading.Lock()
postgresql_engine = create_engine(f"postgresql://{default['UserName']}:{default['PassWord']}@{default['Host']}:{default['Port']}/{default['DBname']}")
Base = declarative_base()
Session = sessionmaker(bind=postgresql_engine)


class NoneskuException(Exception):
    pass


def migration(model):
    with session_scope() as session:
        rows = session.query(model).all()

    for row in rows:
        p = copy.copy(row)
        session = Session()
        try:
            session.add(p)
            session.commit()
        except IntegrityError as e:
            print(row.value)
            session.rollback()

class Product(Base):
    __tablename__ = 'product_master'
    date = Column(Integer)
    name = Column(String)
    asin = Column(String)
    jan = Column(String)
    sku = Column(String, primary_key=True)
    fnsku = Column(String)
    danger_class = Column(String)
    sell_price = Column(Integer)
    cost_price = Column(Integer)

    @property
    def value(self):
        return {
            'date': self.date,
            'name': self.name,
            'asin': self.asin,
            'jan': self.jan,
            'sku': self.sku,
            'fnsku': self.fnsku,
            'danger_class': self.danger_class,
            'sell_price': self.sell_price,
            'cost_price': self.cost_price,
        }

    @classmethod
    def create(cls, date, product_name, asin_code, jan, sku, fnsku, danger_class, sell_price, cost_price):
        product = cls(date=date, name=product_name, asin=asin_code, jan=jan, sku=sku, fnsku=fnsku,
                      danger_class=danger_class, sell_price=sell_price, cost_price=cost_price)
        try:
            with session_scope() as session:
                session.add(product)
            return True
        except IntegrityError:
            return False

    @classmethod
    def sku_get_detail(cls, sku):
        with session_scope() as session:
            product = session.query(cls).filter(cls.sku == sku).first()
            if product is None:
                return None
            return product

    @classmethod
    def update_cost_price(cls, sku, cost_price):
        with session_scope() as session:
            product = session.query(cls).filter(cls.sku == sku).first()
            if product is None:
                return None
            product.cost_price = cost_price
            return True


class Stock(Base):
    __tablename__ = 'stock'
    sku = Column(String, ForeignKey("product_master.sku"), primary_key=True)
    home_stock_count = Column(Integer)
    fba_stock_count = Column(Integer)

    @classmethod
    def add_home_stock(cls, sku: str, home_stock: int):
        product = cls(sku=sku, home_stock=home_stock, fba_stock=0)
        try:
            with session_scope() as session:
                response = session.query(cls).filter(cls.sku == sku).first()
                if not response:
                    session.add(product)
                else:
                    response.home_stock += int(home_stock)
            return True
        except IntegrityError:
            return False

    @classmethod
    def add_fba_stock(cls, sku: str, fba_stock: int):
        product = cls(sku=sku, home_stock=0, fba_stock=fba_stock)
        try:
            with session_scope() as session:
                response = session.query(cls).filter(cls.sku == sku).first()
                if not response:
                    session.add(product)
                else:
                    response.fba_stock += int(fba_stock)
            return True
        except IntegrityError:
            return False

    @classmethod
    def decrease_home_stock(cls, sku: str, home_stock: int):
        try:
            with session_scope() as session:
                response = session.query(cls).filter(cls.sku == sku).first()
                if not response:
                    raise NoneskuException(f'sku not found {sku}')
                else:
                    response.home_stock -= int(home_stock)
            return True
        except NoneskuException:
            return False

    @classmethod
    def initialize_fba_stock(cls):
        try:
            with session_scope() as session:
                response = session.query(cls).all()
                for product in response:
                    product.fba_stock = 0
            return True
        except IntegrityError:
            return False

    @property
    def value(self):
        return {
            'sku': self.sku,
            'home_stock': self.home_stock,
            'fba_stock': self.fba_stock,
        }


class InactiveStock(Base):
    __tablename__ = 'inactivestock'
    sku = Column(String, ForeignKey("product_master.sku"), primary_key=True)
    asin = Column(String)

    @classmethod
    def save(cls, sku, asin):
        product = cls(sku=sku, asin=asin)
        try:
            with session_scope() as session:
                session.add(product)
        except IntegrityError:
            return False

    @classmethod
    def delete(cls):
        with session_scope() as session:
            session.query(cls).delete()

    @classmethod
    def get_asin_cost(cls):
        with session_scope() as session:
            response = session.query(cls, Product, FavoriteProduct).join(Product, cls.asin == Product.ASIN_code)\
                        .join(FavoriteProduct, Product.jan == FavoriteProduct.jan)\
                        .filter(Product.cost_price.isnot(None)).order_by(desc(FavoriteProduct.cost)).all()
            return response

    @property
    def value(self):
        return {
            'sku': self.sku,
            'asin': self.asin,
        }


class FavoriteProduct(Base):
    __tablename__ = 'favoriteproduct'
    url = Column(String, primary_key=True)
    jan = Column(String, primary_key=True)
    cost = Column(Integer)

    @classmethod
    def save(cls, url, jan, cost):
        row = cls(url=url, jan=jan, cost=cost)
        try:
            with session_scope() as session:
                session.add(row)
        except IntegrityError:
            return False

    @classmethod
    def delete(cls):
        with session_scope() as session:
            session.query(cls).delete()

    @property
    def value(self):
        return {
            'url': self.url,
            'jan': self.jan,
            'cost': self.cost,
        }


@contextmanager
def session_scope():
    session = Session()
    session.expire_on_commit = False
    try:
        lock.acquire()
        yield session
        session.commit()
    except Exception as e:
        logger.error(f'action=session_scope error={e}')
        session.rollback()
        raise
    finally:
        session.expire_on_commit = True
        lock.release()


def calc_stock_cost():
    df = pd.read_sql_query('SELECT stock.sku, home_stock, fba_stock, (home_stock+fba_stock)*product_master.cost_price as total \
        from Stock inner join product_master on stock.sku = product_master.sku where home_stock > 0 or fba_stock > 0', postgresql_engine)
    month = datetime.date.today().strftime('%y%m')
    df.to_excel(default['save_path']+month+'.xlsx', index=False)


def init_db():
    Base.metadata.create_all(bind=postgresql_engine)


if __name__ == '__main__':
    init_db()
