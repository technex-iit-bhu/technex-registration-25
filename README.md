# Technex Registration : 2025

Format complete code using : `chmod +x format.sh && ./format.sh`

Following are the routes that need to be designed :
- /register : Register a user in MongoDB backend with the following fields :
    - Name
    - Email
    - Password
    - College
    - City
    - Year
    - Branch
    - Phone Number
    - Referral Code [Once the user enters the referral code, the user who referred the new user will get a notification. Referral is permanent]
- /login : Login a user with email and password and display all details of the user in a dashboard.
            These will be later integrated in main Technex website.
- /update : Update the details of the user with the following fields :
    - Name
    - Email
    - Password (Update only if the user enters the correct old password) [hence, seperate route for password update]] 
    - College
    - City
    - Year
    - Branch
    - Phone Number
- /delete : Delete the user from the database : Delete the user account.
- /forgot : Forgot password route : Send an email to the user with a link to reset the password.
- /google : Google OAuth route : Login with Google OAuth. : [Sign up + Login]
- /github : Github OAuth route : Login with Github OAuth. : [Sign up + Login]
- /events : Display all events of Technex 2025.
- /events/{event} : Display details of a particular event.
