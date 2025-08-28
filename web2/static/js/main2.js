class LoginForm {
    constructor() {
        this.form = document.getElementById('loginForm');
        this.emailField = document.getElementById('email');
        this.passwordField = document.getElementById('password');
        this.loginBtn = document.getElementById('loginBtn');
        this.togglePassword = document.getElementById('togglePassword');
        this.alertContainer = document.getElementById('alertContainer');

        this.initEventListeners();
        this.initPasswordToggle();
    }

    initEventListeners() {
        // Real-time validation
        this.emailField.addEventListener('blur', () => this.validateEmail());
        this.passwordField.addEventListener('blur', () => this.validatePassword());

        // Remove validation on input
        this.emailField.addEventListener('input', () => this.clearFieldValidation(this.emailField));
        this.passwordField.addEventListener('input', () => this.clearFieldValidation(this.passwordField));

    }

    initPasswordToggle() {
        this.togglePassword.addEventListener('click', () => {
            const type = this.passwordField.getAttribute('type') === 'password' ? 'text' : 'password';
            this.passwordField.setAttribute('type', type);

            this.togglePassword.classList.toggle('fa-eye');
            this.togglePassword.classList.toggle('fa-eye-slash');
        });
    }

    validateEmail() {
        const email = this.emailField.value.trim();
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

        if (!email) {
            this.setFieldInvalid(this.emailField, 'Email address is required.');
            return false;
        } else if (!emailRegex.test(email)) {
            this.setFieldInvalid(this.emailField, 'Please enter a valid email address.');
            return false;
        } else {
            this.setFieldValid(this.emailField);
            return true;
        }
    }

    validatePassword() {
        const password = this.passwordField.value;

        if (!password) {
            this.setFieldInvalid(this.passwordField, 'Password is required.');
            return false;
        } else if (password.length < 6) {
            this.setFieldInvalid(this.passwordField, 'Password must be at least 6 characters long.');
            return false;
        } else {
            this.setFieldValid(this.passwordField);
            return true;
        }
    }

    setFieldInvalid(field, message) {
        field.classList.add('is-invalid');
        field.classList.remove('is-valid');
        const feedback = field.nextElementSibling?.nextElementSibling;
        if (feedback && feedback.classList.contains('invalid-feedback')) {
            feedback.textContent = message;
        }
    }

    setFieldValid(field) {
        field.classList.add('is-valid');
        field.classList.remove('is-invalid');
    }

    clearFieldValidation(field) {
        field.classList.remove('is-valid', 'is-invalid');
    }

    showAlert(message, type = 'danger') {
        const alertHtml = `
                    <div class="alert alert-${type} alert-dismissible fade show" role="alert">
                        <i class="fas fa-${type === 'success' ? 'check-circle' : 'exclamation-triangle'}"></i>
                        ${message}
                        <button type="button" class="btn-close" data-bs-dismiss="alert"></button>
                    </div>
                `;
        this.alertContainer.innerHTML = alertHtml;
    }

    setLoading(loading) {
        if (loading) {
            this.loginBtn.classList.add('loading');
            this.loginBtn.disabled = true;
        } else {
            this.loginBtn.classList.remove('loading');
            this.loginBtn.disabled = false;
        }
    }


}

// Initialize the login form when the page loads
document.addEventListener('DOMContentLoaded', () => {
    new LoginForm();
});