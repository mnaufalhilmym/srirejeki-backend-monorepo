package auth

func Module(a *Auth) {
	getUserSession(a)
	postSignUp(a)
	postVerifySignUp(a)
	postSignIn(a)
	postRequestResetPassword(a)
	postVerifyRequestResetPassword(a)
	patchResetPassword(a)
	getSignOut(a)
}
