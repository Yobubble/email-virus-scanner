### TODO
* [ ] virus-scanner package (Boom)
<!-- Ref: https://mailpit.axllent.org/docs/api-v1/websocket/ -->
* [ ] setup websocket for email notification (BB)

### Instruction (in progress)
<!-- Ref: https://mailpit.axllent.org/docs/install/testing/ -->
1. run ```docker compose up -d```
2. Test sending email  
    - using command ```mailpit sendmail < email.txt``` (attachments dont'work)
    - using mailpit's REST API "http://localhost:8025/api/v1/send"
3. Go to http://localhost:8025 to view the mailpit's UI and the email sent 

![image](./images/overview.png)