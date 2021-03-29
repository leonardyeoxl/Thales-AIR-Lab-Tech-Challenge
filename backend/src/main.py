from STARsSIDs import RetrieveSTARsSIDs
import json
import pika

connection = pika.BlockingConnection(
    pika.ConnectionParameters(host='mq', port='5672'))

channel = connection.channel()

channel.queue_declare(queue='rpc_queue')

def on_request(ch, method, props, body):
    response = RetrieveSTARsSIDs()

    ch.basic_publish(exchange='',
                     routing_key=props.reply_to,
                     properties=pika.BasicProperties(correlation_id = \
                                                         props.correlation_id),
                     body=json.dumps(response).encode('utf8'))
    ch.basic_ack(delivery_tag=method.delivery_tag)

if __name__ == "__main__":
    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(queue='rpc_queue', on_message_callback=on_request)

    print(" [x] Awaiting RPC requests")
    channel.start_consuming()