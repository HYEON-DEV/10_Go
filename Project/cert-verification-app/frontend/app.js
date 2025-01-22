const apiUrl = 'http://localhost:8080/api/certificates'; // Update with your backend API URL

document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('certForm');
    const resultDiv = document.getElementById('result');

    form.addEventListener('submit', async (event) => {
        event.preventDefault();
        
        const formData = new FormData(form);
        const certificateData = Object.fromEntries(formData.entries());

        try {
            const response = await fetch(apiUrl, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(certificateData),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const result = await response.json();
            resultDiv.textContent = `Verification Result: ${result.message}`;
        } catch (error) {
            resultDiv.textContent = `Error: ${error.message}`;
        }
    });
});