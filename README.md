reCAPTCHA v3 Demo with Go and Tailwind CSS
A production-ready Go application demonstrating Google reCAPTCHA v3 integration with a sleek, user-friendly frontend styled using Tailwind CSS. This project showcases secure form verification with a unique feature: a dynamic visualization of the reCAPTCHA score to indicate whether a user is likely human or a bot.

Features

reCAPTCHA v3 Integration: Seamlessly verifies user interactions using Google reCAPTCHA v3 with a Go backend (net/http).
Score Visualization: Displays the reCAPTCHA score (0.0 to 1.0) on the success page with an animated progress bar, making it educational and engaging.
Responsive UI: Styled with Tailwind CSS for a modern, clean, and mobile-friendly design.
Error Handling: User-friendly error messages displayed on the form page if verification fails.
Environment Configuration: Uses godotenv to securely manage reCAPTCHA keys and server port.
Dynamic Templating: Injects the reCAPTCHA Site Key into the frontend using Go’s html/template.

Prerequisites

Go: Version 1.18 or higher.
Google reCAPTCHA v3 Keys:
Register at Google reCAPTCHA Admin Console.
Select reCAPTCHA v3 and add localhost for testing.
Obtain a Site Key and Secret Key.


Node.js (optional): If you want to compile Tailwind CSS for production instead of using the CDN.

Setup Instructions

Clone the Repository:
git clone https://github.com/yourusername/recaptcha-demo.git
cd recaptcha-demo


Install Dependencies:
go mod tidy
go get github.com/joho/godotenv


Configure Environment Variables:Create a .env file in the project root with the following:
RECAPTCHA_SECRET=your_secret_key
RECAPTCHA_SITE_KEY=your_site_key
PORT=8080

Replace your_secret_key and your_site_key with your reCAPTCHA v3 keys.

Run the Application:
go run cmd/server/main.go


Access the Application:Open http://localhost:8080 in your browser. Submit the form to test reCAPTCHA verification and see the score visualization on success.


Project Structure
recaptcha-demo/
├── cmd/
│   └── server/
│       └── main.go           # Entry point for the server
├── internal/
│   ├── config/
│   │   └── config.go        # Environment variable loading
│   ├── handlers/
│   │   └── handlers.go      # HTTP handlers for form and verification
│   └── recaptcha/
│       └── recaptcha.go     # reCAPTCHA v3 verification logic
├── static/
│   ├── index.html           # Form page with reCAPTCHA v3
│   └── success.html         # Success page with score visualization
├── .env                     # Environment variables (not committed)
├── .gitignore               # Git ignore file
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies checksum
└── README.md                # This file

Usage

Form Submission:

Visit http://localhost:8080.
Click the "Submit" button to trigger reCAPTCHA v3 verification.
The verification runs invisibly, sending a token to the backend.


Success Page:

On successful verification (score ≥ 0.5, action = "submit"), you’ll see a success page with the reCAPTCHA score displayed as a progress bar.
Click "Back to Form" to try again.


Error Handling:

If verification fails (e.g., low score or invalid token), you’ll be redirected to the form page with an error message.

Production Considerations

HTTPS: Use a reverse proxy (e.g., Nginx) or http.ListenAndServeTLS with an SSL certificate (e.g., Let’s Encrypt).
Tailwind CSS: Compile Tailwind CSS to a static file using the Tailwind CLI for better performance instead of the CDN.
Logging: Add structured logging (e.g., logrus) to monitor verification attempts.
Security: Implement CSRF protection (e.g., github.com/gorilla/csrf) for form submissions.
Threshold Tuning: Adjust the reCAPTCHA score threshold (currently 0.5) based on your use case.

Why This Project Stands Out
This demo goes beyond a basic reCAPTCHA integration by:

Visualizing the reCAPTCHA score with an animated progress bar, making it educational and engaging.
Using a clean, modular Go backend with net/http and godotenv.
Featuring a polished, responsive UI with Tailwind CSS.
Providing robust error handling and dynamic templating for a production-ready experience.

Future Enhancements

Dark Mode: Add a toggle for dark mode using Tailwind’s dark: classes.
Analytics Dashboard: Create an admin page to log and display verification attempts.
PWA Support: Add a manifest and service worker to make the app installable.
Multi-Action Support: Handle different reCAPTCHA actions (e.g., login, signup) with custom thresholds.


