from promise import Promise
from converse import Converse
from send_message import SendMessage


def HandleMessageThread(message_type, send_id, text=''):
  try:
    if text is None:
      text = ''
    SendMessage(message_type, send_id, Converse(text, message_type=message_type, send_id=send_id))
  except Exception as e:
    print(e)


def HandleMessage(message_type, send_id, text=''):
  Promise(HandleMessageThread, (message_type, send_id, text)).start()
