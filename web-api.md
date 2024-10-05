# Service Workder API
Web-App,browser ---- Service-Worker --------network

perperties:
    1. service workder is event-driven workder
    2. NO DOM access
    3. runs on different thread NOTE: means not in main thread
    4. non-blocking
    5. fully async
    6. can not  use Synchronous API 
        sync API List:
            1. XHR
            2. WebStorage
    7. can not import js module dynamically, only static import  allowed.
    8. only uses HTTPS but can HTTP for testing purpose
    9. In Some Browser cannot use in PRIVATE(incognito) Browsing mode






use Cases:
    1. create offline experiences 
    2. Background Data syncronigation
    3. Respoding resource request from other origin
    4. 
 uses Web API:
    1. Notification API
    2. Backgroudn Sync AP
