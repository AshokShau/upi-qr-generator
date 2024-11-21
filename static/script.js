function openUPIApp(vpa, amount) {
    const upiUrl = `upi://pay?pa=${encodeURIComponent(vpa)}&am=${encodeURIComponent(amount)}&cu=INR`;
    window.location.href = upiUrl;
}
