from typing import List
import logging
import azure.functions as func
import pyodbc
import os
import json
import datetime
from pytz import timezone

server = os.environ['SQL_SERVER']
database = os.environ['SQL_DATABASE']
username = os.environ['SQL_USERNAME']
password = os.environ['SQL_PASSWORD']
driver = '{ODBC Driver 17 for SQL Server}'


def main(events: List[func.EventHubEvent], context: func.Context):
    start_time = datetime.datetime.now(timezone('UTC'))

    logging.info("--- Func start: EventLength=" + str(len(events)) + ", InvocationID=" + context.invocation_id + ", Now=" + start_time.strftime("%Y-%m-%d %H:%M:%S"))

    with pyodbc.connect('DRIVER='+driver+';SERVER='+server+';PORT=1433;DATABASE='+database+';UID='+username+';PWD=' + password) as conn:
        with conn.cursor() as cursor:
            for event in events:
                logging.info(f'Function triggered to process a message: {event.get_body().decode()}, EnqueuedTimeUtc = {event.enqueued_time}, SequenceNumber = {event.sequence_number}, Offset = {event.offset}, PartitionKey = {event.partition_key}')
                d = json.loads(event.get_body().decode('utf-8'))
                cursor.execute(
                    "INSERT INTO fromfunctions(id,process_datetime,value_a,value_b,eh_enqueuedtimeutc,eh_offset,eh_sequenceno, sql_insert_datetime, func_invocationid) VALUES (?,?,?,?,?,?,?,?,?)",
                    d["id"], d["process_datetime"], d["value_a"], d["value_b"],
                    event.enqueued_time, event.offset, event.sequence_number, datetime.datetime.now(timezone('UTC')).strftime("%Y-%m-%d %H:%M:%S"),
                    context.invocation_id
                )

    end_time = datetime.datetime.now(timezone('UTC'))
    duration = (end_time - start_time)
    logging.info("--- Func end: EventLength=" + str(len(events)) + ", InvocationID=" + context.invocation_id + ", Now=" + end_time.strftime("%Y-%m-%d %H:%M:%S") + ", Duration(ms)=" + str((duration.seconds + duration.microseconds / 1000000) * 1000))