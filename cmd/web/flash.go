package main

import "net/http"

const MsgRecipeCreated = "Recipe successfully created!!!"

const MsgUserCreated = "User successfully created!!! Please login using your credential"

const MsgUserAuthenticeted = "Login successfully"

func (app *application) SetFlashMsg(r *http.Request, msg string) {
	app.sessionManager.Put(r.Context(), "flash", msg)
}
