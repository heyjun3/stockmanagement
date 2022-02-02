from multiprocessing import Queue
import logging.config

import pandas as pd

from mws import api

# logging.config.fileConfig(r'', disable_existing_loggers=False)
logger = logging.getLogger(__name__)


def get_competitive_pricing_for_asin_worker(filename: str, que: Queue, manager):
    logger.info('action=get_competitive_pricing_for_asin_worker status=run')
    data = []

    amazon = api.AmazonClient()
    while True:
        products = que.get()
        # print(products)
        # print(data)
        if products is None:
            logger.info('action=get_competitive_pricing_for_asin_worker status=done')
            df = pd.DataFrame(data=data, columns=['asin', 'price']).astype({'price': int})
            manager['price_df'] = df
            return
        items = amazon.get_competitive_pricing_for_asin(products, filename)
        data.extend(items)


def get_lowest_priced_offer_listtings_for_asin_worker(filename: str, que: Queue, manager):
    logger.info('action=get_competitive_pricing_for_asin_worker status=run')
    data = []

    amazon = api.AmazonClient()
    while True:
        products = que.get()
        # print(products)
        # print(data)
        if products is None:
            logger.info('action=get_competitive_pricing_for_asin_worker status=done')
            df = pd.DataFrame(data=data, columns=['asin', 'price']).astype({'price': int})
            manager['price_df'] = df
            return
        items = amazon.get_lowest_priced_offer_listtings_for_asin(products, filename)
        data.extend(items)


def get_fee_my_fees_estimate_worker(filename: str, que: Queue, manager):
    logger.info('action=get_fee_my_fees_estimate_worker status=run')
    data = []

    amazon = api.AmazonClient()
    while True:
        products = que.get()
        if products is None:
            logger.info('action=get_fee_my_fees_estimate_worker status=done')
            df = pd.DataFrame(data=data, columns=['asin', 'fee_rate', 'ship_fee']) \
                .astype({'fee_rate': float, 'ship_fee': float})
            manager['fees_df'] = df
            return
        items = amazon.get_fee_my_fees_estimate(products, filename)
        data.extend(items)

