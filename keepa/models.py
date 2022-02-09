import logging
import threading
import datetime

from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, Date
from sqlalchemy import JSON
from sqlalchemy.orm import sessionmaker
from sqlalchemy.exc import IntegrityError
from contextlib import contextmanager
import pandas as pd
import numpy as np

import settings


logger = logging.getLogger('sqlalchemy.engine')
logger.setLevel(logging.WARNING)
lock = threading.Lock()
postgresql_engine = create_engine(settings.DB_URL)
Base = declarative_base()
NetseaBase = declarative_base()
Session = sessionmaker(bind=postgresql_engine)


class KeepaProducts(Base):
    __tablename__ = 'keepa_products'
    asin = Column(String, primary_key=True)
    sales_drops_90 = Column(Integer)
    created = Column(Date, default=datetime.date.today)
    modified = Column(Date, default=datetime.date.today)
    price_data = Column(JSON)
    rank_data = Column(JSON)

    @classmethod
    def create(cls, asin, drops ,price_data, rank_data):
        shop = cls(asin=asin, sales_drops_90=drops, price_data=price_data, rank_data=rank_data)
        try:
            with session_scope() as session:
                session.add(shop)
            return True
        except IntegrityError:
            return False

    @classmethod
    def object_get_db_asin(cls, asin, delay=30):
        with session_scope() as session:
            delay_date = datetime.date.today() - datetime.timedelta(days=delay)
            product = session.query(cls).filter(cls.asin == asin, cls.modified >= delay_date).first()
            if product is None:
                return None
            return product

    @classmethod
    def update_or_insert(cls, asin, drops, price_data, rank_data):
        with session_scope() as session:
            product = session.query(cls).filter(cls.asin == asin).first()
            if product is None:
                keepa_product = cls(asin=asin, sales_drops_90=drops, price_data=price_data, rank_data=rank_data)
                try:
                    session.add(keepa_product)
                except IntegrityError as ex:
                    logger.error(ex)
                    logger.error(keepa_product.value)
            else:
                product.sales_drops_90 = drops
                product.modified = datetime.date.today()
                product.price_data = price_data
                product.rank_data = rank_data
            return True
    
    @classmethod
    def get_product_price_data_is_None(cls, get_product_num: int = 100):
        with session_scope() as session:
            products = session.query(cls).filter(cls.price_data == None, cls.rank_data == None).limit(get_product_num).all()
            if products:
                return products
            else:
                return None

    @classmethod
    def get_keepa_product(cls, asin: str):
        with session_scope() as session:
            product = session.query(cls).filter(cls.asin == asin).first()
            if product:
                return product
            else:
                return None

    
    @property
    def render_price_rank_data_list(self):
        rank_dict = {convert_keepa_time_to_datetime_date(int(k)): v for k, v in self.rank_data.items()}
        price_dict = {convert_keepa_time_to_datetime_date(int(k)): v for k, v in self.price_data.items()}

        rank_df = pd.DataFrame(data=list(rank_dict.items()), columns=['date', 'rank']).astype({'rank': int})
        price_df = pd.DataFrame(data=list(price_dict.items()), columns=['date', 'price']).astype({'price': int})

        df = pd.merge(rank_df, price_df, on='date', how='outer')
        df = df.replace(-1.0, np.nan)
        df = df.fillna(method='ffill')
        df = df.fillna(method='bfill')
        delay = datetime.datetime.now().date() - datetime.timedelta(days=90)
        df = df[df['date'] > delay]
        df = df.sort_values('date', ascending=True)
        products = (df['date'].to_list(), df['rank'].to_list(), df['price'].to_list())

        return products

    @property
    def value(self):
        return {
            'asin': self.asin,
            'sales_drop_90': self.sales_drops_90,
            'created': self.created,
            'modified': self.modified,
            'price_data': self.price_data,
            'rank_data': self.rank_data,
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


def init_db():
    Base.metadata.create_all(bind=postgresql_engine)


def convert_keepa_time_to_datetime_date(keepa_time: int):
    unix_time = (keepa_time + 21564000) * 60
    date_time = datetime.datetime.fromtimestamp(unix_time)
    return date_time.date()
