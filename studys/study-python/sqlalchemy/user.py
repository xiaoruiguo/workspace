#coding=utf-8
import sqlalchemy

from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine
from sqlalchemy import Column,Integer,String,ForeignKey
from sqlalchemy.orm import sessionmaker,relationship,backref

Base = declarative_base()

class User(Base):             #表users
    __tablename__ = 'users'   #指定表明

    #创建对应的列
    id = Column(Integer, primary_key=True)
    name = Column(String(50))  #有的数据库需要设置string长度,如mysql
    fullname = Column(String(50))
    password = Column(String(50))

    def __repr__(self):
        return "<User(name='%s',fullname='%s',password='%s')>"%(
               self.name, self.fullname, self.password)

class Address(Base):
    __tablename__ ='addresses'
    id = Column(Integer, primary_key=True)
    email_address = Column(String(50), nullable=False)

    #外键:索引到users.id
    user_id = Column(Integer, ForeignKey('users.id'))

    #在Address中增加user属性，该属性的值是address的外键对应的user表中中记录
    #同时在User中增加addresses属性，该属性的值是user对应的addresses表中的记录
    user = relationship("User", backref=backref('addresses', order_by=id))

    def __repr__(self):
        return "<Addresses(email_address='%s')>" % self.email_address

#连接到数据库
#engine = create_engine('sqlite:///:memory:',echo=False)
engine = create_engine('mysql://root:@localhost/sqlalchemy',echo=False)

#创建数据库表
Base.metadata.create_all(engine)

#获取一个数据库会话
Session = sessionmaker(bind=engine)
session = Session()

#生成一条记录
user_a = User(name='a',fullname='aaaaaaaaa',password='adfadfdafadf')
session.add(user_a)  #现在user_a缓存在session中，还没有提交到数据库中
session.commit()     #将session中更新提交到数据库中
                     #如果不手动提交，更新会在执行查询操作时自动提交

#查询，验证记录已经提交到数据库
rows=session.query(User).all()
print rows

#修改记录的值
user_a.password='1234aaaaaaaaaaaaa'
session.commit()     #将session中更新提交到数据库中

#查询, 验证记录的修改已经提交到数据库
rows=session.query(User).all()
print rows

#在查询结果中修改记录的值
for row in rows:
    row.name='bbbbbbaaaaa'
    session.commit()

#查询，验证在查询结果中的修改是否提交到数据库
rows=session.query(User).all()
print rows

addr1 = Address(email_address="xxx1@xx.com")
addr2 = Address(email_address="xxx2@xx.com")
user_a.addresses=[addr1, addr2]
user_b = User(name='b',fullname='bbbb',password='adfadfdafadf')
user_b.addresses=[addr2]  #addr2的外键会被自动修改为user_b的id
session.commit()

#查询，验证addresses中的记录
rows = session.query(Address).all()
print rows
