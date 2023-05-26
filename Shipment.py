#EAI
#Developing API - Python side - Shipment module
from flask import Flask
from flask import request,jsonify
from flask_cors import CORS
import mysql.connector
import requests
mydb = mysql.connector.connect(
  host="localhost",
  user="youruser",
  password="yourpassword",
  database="eaiCourier"
)




app=Flask(__name__)
CORS(app)

@app.route("/",methods=['POST'])
def T():
    data = request.get_json()
    print(data)
    try:
        mycursor = mydb.cursor()
        sql = "INSERT INTO shipmentCourier (pickupAddress,DeliverAddress) VALUES (%s,%s)"
        val = (data['PickupLocation'],data['DeliveryLocation'])
        mycursor.execute(sql, val)
        cus_id =  mycursor.lastrowid
        mydb.commit()
        mycursor.close()
        status="success"
    except Exception:
        status="some error occured"
    return jsonify(status),200
    
app.run()
