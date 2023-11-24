# Makoto/v3

## Frontend
### Main page
- /: Welcome page
- /me: Current user page
- /me/edit: Current user edit page
- /me/friends: User's friends
______
- /: logo + text `Makoto`
- /me: user's picture, username,　join date, intergration (providers like Discord, Spotify) | edit button | microservice stat
- /me/edit:
1. update username, description, picture, preferred color, languages
2. "who can see my user page?", "who can send me friend request?", "who can see my micro stats?", "who can message me?", "who can see my user picture?"
3. integrations' settings, email settings (redirect)


### Authentication
- /: login
- /signup
- /confirm-email
- /password/change
- /password/recover
______
- SignIn via OAuth2/Credentials
- SignUp via credentials
- Confirm email when via credentials
- Change/Recover password for credentials
- Max SignIn tries ~ 5 -> timeout
- Error messages
- Clarify whether username/email is reserved
- Password strongness validation

## Backend
### Authentication
- SignIn
- SignUp
- SignInOAuth
- SendConfirmationMail (no)
- ConfirmMailByCode (tends to be hashed username -> compare hash maybe?)
- ProceedToUpdatePassword ( using old password )
- UpdatePassword ( via some sort of generated and stored in db (or redis) token )
- SendRecoverPasswordMail (no)
- SubmitNewPasswordByRecoverdCode ( same sort token as for UpdatePassword)
- IsEmailAvailable
- IsUsernameAvailable
_______

### Mail Service
- SendEmail

### CDN

### Server Registry && Discovery
- RegisterService
- UnregisterService
_____
- Redis ( as key-value pairs )

###
