### TODO
* [ ] virus-scanner package (Boom)
<!-- Ref: https://mailpit.axllent.org/docs/api-v1/websocket/ -->
* [X] ~~*setup websocket for email notification (BB)*~~ [2025-03-15]
* [ ] GetEmailFromID 


### Instruction (in progress)
<!-- Ref: https://mailpit.axllent.org/docs/install/testing/ -->
1. run ```docker compose up -d```
2. Test sending email  
    - using command ```mailpit sendmail < email.txt``` (attachments dont'work)
    - Or, I already implemented one: ```go run main.go sendmail```, this will send mock email to mailpit directly using its REST API (http://localhost:8025/api/v1/send)
3. Go to http://localhost:8025 to view the mailpit's UI and the email sent 

![image](./images/overview.png)