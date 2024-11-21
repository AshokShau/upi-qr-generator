package str

const Index = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UPI QR Code API</title>
    <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;500;700&display=swap" rel="stylesheet">
    <style>
        body {
            font-family: 'Roboto', sans-serif;
            margin: 0;
            padding: 0;
            background: #f4f4f9;
            color: #333;
        }
        .container {
            max-width: 800px;
            margin: 30px auto;
            padding: 20px;
            background: #fff;
            border-radius: 8px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
        }
        h1 {
            color: #4CAF50;
            text-align: center;
        }
        p {
            line-height: 1.6;
        }
        pre {
            background: #f4f4f4;
            padding: 15px;
            border-radius: 5px;
            overflow-x: auto;
        }
        .footer {
            text-align: center;
            margin-top: 20px;
            font-size: 0.9em;
            color: #777;
        }
        .footer a {
            color: #4CAF50;
            text-decoration: none;
        }
    </style>
</head>
<body>
<div class="container">
    <h1>UPI QR Code API</h1>
    <p>Welcome to the UPI QR Code API! This API allows you to generate QR codes for UPI payments and render a simple payment page. Below is the information on how to use the endpoints.</p>

    <h2>Endpoints</h2>
    <ul>
        <li><strong>Generate QR Code</strong> - <code>GET /qr</code></li>
        <p>Parameters:
        <ul>
            <li><code>vpa</code> (string) - UPI ID (required)</li>
            <li><code>name</code> (string) - Payee name (optional)</li>
            <li><code>amount</code> (float) - Payment amount (required)</li>
        </ul>
        </p>
        <p><strong>Example:</strong> <code>/qr?vpa=example@upi&name=John+Doe&amount=100.50</code></p>

        <li><strong>Payment Page</strong> - <code>GET /pay</code></li>
        <p>Parameters:
        <ul>
            <li><code>vpa</code> (string) - UPI ID (required)</li>
            <li><code>name</code> (string) - Payee name (optional)</li>
            <li><code>amount</code> (float) - Payment amount (required)</li>
        </ul>
        </p>
        <p><strong>Example:</strong> <code>/pay?vpa=example@upi&name=John+Doe&amount=100.50</code></p>
    </ul>

    <h2>Source Code</h2>
    <p>The source code for this project is available on GitHub. Feel free to explore, use, and contribute!</p>
    <p><a href="https://github.com/AshokShau/upi-qr-generator" target="_blank">https://github.com/AshokShau/</a></p>

    <div class="footer">
        <p>© 2024 by <a href="https://github.com/AshokShau" target="_blank">AshokShau</a>. All Rights Reserved.</p>
    </div>
</div>
</body>
</html>
`

const paymentHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta content="width=device-width, initial-scale=1.0" name="viewport">
        <title>Secure Payment Gateway</title>
        <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css" rel="stylesheet">
        <link href="https://fonts.googleapis.com/css2?family=Roboto:wght@400;500;700&display=swap" rel="stylesheet">
        <style>
            /* styles.css */
            body {
                margin: 0;
                font-family: 'Roboto', sans-serif;
                background: url('https://images.pexels.com/photos/691668/pexels-photo-691668.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2') no-repeat center center fixed;
                background-size: cover;
                color: #333;
            }

            .background {
                display: flex;
                justify-content: center;
                align-items: center;
                height: 100vh;
                backdrop-filter: blur(5px);
            }

            .payment-container {
                background-color: rgba(255, 255, 255, 0.9);
                border-radius: 12px;
                box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
                padding: 20px 30px;
                max-width: 400px;
                text-align: center;
            }

            .header {
                display: flex;
                flex-direction: column;
                align-items: center;
                margin-bottom: 20px;
            }

            .header .logo {
                width: 80px;
                height: auto;
                margin-bottom: 10px;
                border-radius: 18%;
                border: 2px solid #4CAF50;
            }

            .header h1 {
                font-size: 24px;
                font-weight: 700;
                color: #4CAF50;
                margin: 0;
            }

            .qr-container {
                margin: 20px 0;
                text-align: center;
            }

            .qr-container img {
                width: 200px;
                height: 200px;
                border: 5px solid #4CAF50;
                border-radius: 12px;
                object-fit: contain;
            }

            .upi-button {
                background-color: #4CAF50;
                color: white;
                border: none;
                border-radius: 5px;
                padding: 10px 15px;
                font-size: 16px;
                cursor: pointer;
                display: inline-flex;
                align-items: center;
                justify-content: center;
                gap: 10px;
                transition: background-color 0.3s;
            }

            .upi-button:hover {
                background-color: #388E3C;
            }

            .details {
                text-align: left;
                margin-top: 15px;
                font-size: 14px;
            }

            .details p {
                margin: 10px 0;
                display: flex;
                align-items: center;
                gap: 8px;
            }

            .details i {
                color: #4CAF50;
            }
        </style>
    </head>

    <body>
    <div class="background">
        <div class="payment-container">
            <div class="header">
                <img alt="Gateway Logo" class="logo" src="https://download.logo.wine/logo/Amazon_Pay/Amazon_Pay-Logo.wine.png">
                <h1>Secure Payment</h1>
            </div>

            <p>Pay <strong>₹{{.Amount}}</strong> to <strong>{{.Name}}</strong></p>

            <div class="qr-container">
                <img alt="UPI QR Code" src="/qr?vpa={{.VPA}}&name={{.Name}}&amount={{.Amount}}"/>
            </div>

            <button class="upi-button" onclick="openUPIApp('{{.VPA}}', '{{.Amount}}')">
                <i class="fas fa-mobile-alt"></i> Pay with UPI App
            </button>

            <div class="details">
                <p><i class="fas fa-user"></i> <strong>Payee:</strong> {{.Name}}</p>
                <p><i class="fas fa-wallet"></i> <strong>UPI ID:</strong> {{.VPA}}</p>
                <p><i class="fas fa-rupee-sign"></i> <strong>Amount:</strong> ₹{{.Amount}}</p>
            </div>
        </div>
    </div>
</body>
<script>
    function openUPIApp(vpa, amount) {
        const upiUrl = 'upi://pay?pa=' + encodeURIComponent(vpa) + '&am=' + encodeURIComponent(amount) + '&cu=INR';
        window.location.href = upiUrl;
    }
</script>
</html>
`
