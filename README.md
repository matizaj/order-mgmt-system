# order-mgmt-system

![plot](order-mgmt-sys.png)<br/>
Order Service: <br/>
    - validate order details<br/>
    - talk with stock service<br/>
    - crud of orders<br/>
    - init payment flow<br/>
Stock Service:<br/>
    - handle stock<br/>
    - validate order quantity<br/>    
Menu Service:<br/>
    - store items as menu<br/>
Payment Service:<br/>
    - init payment with 3rd poarty provider<br/>
    - produce order paid/cancelled events<br/>
Kitchen Service:<br/>
    - simulate long running process<br/>