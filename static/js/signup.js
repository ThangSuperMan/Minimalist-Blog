const password = document.querySelector("#password")
const conffirmPassword = document.querySelector("#confirm_password")
const validetPasswordDiv = document.querySelector(".validate-password")
let currentPassword, currentConfirmPassword

const handleTypingInputPassword = (e) => {
	console.log("handleTypingInputPassword ")
	currentPassword = e.target.value
}

const handleTypingInputConfirmPassword = (e) => {
	console.log("handleTypingInputConfirmPassword")
	currentConfirmPassword = e.target.value

	// Last field
	if (currentPassword === currentConfirmPassword) {
		console.log("currentPassword === currentConfirmPassword")
		const span = document.createElement("span")
		span.textContent = "Your password is matching with eatch other"
		validetPasswordDiv.appendChild(span)
	} else {
		console.log("currentPassword !== currentConfirmPassword")
		const span = document.createElement("span")
		span.textContent = "Your password is not matching, please typing your password and make sure everything is te smae"
		validetPasswordDiv.appendChild(span)
	}
}

password.addEventListener("change", handleTypingInputPassword)
conffirmPassword.addEventListener("change", handleTypingInputConfirmPassword)
