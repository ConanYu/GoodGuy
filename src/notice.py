import os, threading
from promise import Promise

notice_id = set()
tmp_file_path = os.path.join(os.path.dirname(os.path.dirname(os.path.abspath(__file__))), 'notice.tmp')
notice_id_lock = threading.Lock()


# 从文件中读取数据
def LoadFromNoticeId():
    global tmp_file_path, notice_id, notice_id_lock
    if not os.path.exists(tmp_file_path):
        return
    with notice_id_lock:
        # 重制
        notice_id = set()
        # 读取数据
        with open(tmp_file_path, 'r', encoding='utf-8') as f:
            notice_ids = f.read().split(',')
            for a_notice_id in notice_ids:
                notice_id.add(a_notice_id)


# 一启动就读文件
LoadFromNoticeId()


# 异步更新到文件中 有极小概率有更新失败的情况
def UpdateNoticeId():
    global tmp_file_path, notice_id, notice_id_lock
    def UpdateNoticeIdInner():
        with notice_id_lock:
            with open(tmp_file_path, 'w', encoding='utf-8') as f:
                s = ''
                for a_notice_id in notice_id:
                    if s != '':
                        s += ','
                    s += a_notice_id
                f.write(s)
    # 异步更新
    Promise(UpdateNoticeIdInner).start()


# 增加notice_id的api
def AddNoticeId(message_type, send_id):
    notice_id.add(message_type + "|" + send_id)
    UpdateNoticeId()


# 删除notice_id的api
def RemoveNoticeId(message_type, send_id):
    notice_id.remove(message_type + "|" + send_id)
    UpdateNoticeId()