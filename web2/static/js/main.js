class PasswordValidator {
    constructor() {
        this.passwordInput = document.getElementById('password');
        this.confirmPasswordInput = document.getElementById('confirmPassword');
        this.strengthBar = document.getElementById('strengthBar');
        this.strengthContainer = document.getElementById('strengthContainer');
        this.strengthText = document.getElementById('strengthText');
        this.requirementsList = document.getElementById('requirementsList');
        this.matchIndicator = document.getElementById('matchIndicator');
        this.form = document.getElementById('registrationForm');

        this.requirements = {
            length: document.getElementById('lengthReq'),
            upper: document.getElementById('upperReq'),
            lower: document.getElementById('lowerReq'),
            number: document.getElementById('numberReq'),
            special: document.getElementById('specialReq')
        };

        this.initializeEventListeners();
    }

    initializeEventListeners() {
        this.passwordInput.addEventListener('input', () => {
            this.showStrengthIndicator();
            this.checkPasswordStrength();
        });

        this.passwordInput.addEventListener('focus', () => {
            this.showStrengthIndicator();
        });

        this.confirmPasswordInput.addEventListener('input', () => this.checkPasswordMatch());
        this.form.addEventListener('submit', (e) => this.handleSubmit(e));

        // Real-time validation for all inputs
        const inputs = this.form.querySelectorAll('input[required]');
        inputs.forEach(input => {
            input.addEventListener('input', () => {
                if (input.value.trim() !== '') {
                    this.validateField(input);
                } else {
                    this.hideError(input);
                }
            });

            input.addEventListener('blur', () => {
                if (input.value.trim() !== '') {
                    this.validateField(input);
                }
            });
        });
    }

    showStrengthIndicator() {
        this.strengthContainer.classList.add('show');
    }

    showError(input, message = null) {
        const errorElement = document.getElementById(input.id + 'Error');
        if (errorElement) {
            if (message) {
                errorElement.textContent = message;
            }
            errorElement.classList.add('show');
            input.classList.add('is-invalid');
        }
    }

    hideError(input) {
        const errorElement = document.getElementById(input.id + 'Error');
        if (errorElement) {
            errorElement.classList.remove('show');
            input.classList.remove('is-invalid', 'is-valid');
        }
    }

    checkPasswordStrength() {
        const password = this.passwordInput.value;

        if (password === '') {
            this.strengthContainer.classList.remove('show');
            this.hideError(this.passwordInput);
            return;
        }

        const checks = {
            length: password.length >= 8,
            upper: /[A-Z]/.test(password),
            lower: /[a-z]/.test(password),
            number: /\d/.test(password),
            special: /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password)
        };

        // Update requirement indicators
        Object.keys(checks).forEach(key => {
            const requirement = this.requirements[key];
            const icon = requirement.querySelector('i');

            if (checks[key]) {
                requirement.classList.remove('unmet');
                requirement.classList.add('met');
                icon.className = 'fas fa-check';
            } else {
                requirement.classList.remove('met');
                requirement.classList.add('unmet');
                icon.className = 'fas fa-times';
            }
        });

        // Calculate strength
        const score = Object.values(checks).filter(Boolean).length;
        this.updateStrengthBar(score);

        // Show missing requirements
        this.requirementsList.classList.add('show');

        // Check password match when password changes
        if (this.confirmPasswordInput.value) {
            this.checkPasswordMatch();
        }
    }

    updateStrengthBar(score) {
        this.strengthBar.className = 'progress-bar';
        let strengthText = '';
        let width = 0;

        switch (score) {
            case 0:
            case 1:
                this.strengthBar.classList.add('strength-weak');
                strengthText = 'Very Weak';
                width = 20;
                break;
            case 2:
                this.strengthBar.classList.add('strength-fair');
                strengthText = 'Weak';
                width = 40;
                break;
            case 3:
                this.strengthBar.classList.add('strength-fair');
                strengthText = 'Fair';
                width = 60;
                break;
            case 4:
                this.strengthBar.classList.add('strength-good');
                strengthText = 'Good';
                width = 80;
                break;
            case 5:
                this.strengthBar.classList.add('strength-strong');
                strengthText = 'Strong';
                width = 100;
                break;
        }

        this.strengthBar.style.width = width + '%';
        this.strengthText.textContent = `Password Strength: ${strengthText}`;
        this.strengthText.className = `strength-text ${this.strengthBar.classList[1]}`;
    }

    checkPasswordMatch() {
        const password = this.passwordInput.value;
        const confirmPassword = this.confirmPasswordInput.value;

        if (confirmPassword === '') {
            this.matchIndicator.innerHTML = '';
            this.hideError(this.confirmPasswordInput);
            return;
        }

        if (password === confirmPassword) {
            this.matchIndicator.innerHTML = '<i class="fas fa-check text-success"></i>';
            this.hideError(this.confirmPasswordInput);
            this.confirmPasswordInput.classList.add('is-valid');
        } else {
            this.matchIndicator.innerHTML = '<i class="fas fa-times text-danger"></i>';
            this.showError(this.confirmPasswordInput, 'Passwords do not match');
        }
    }

    validateField(field) {
        if (field.type === 'email') {
            const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
            if (!emailRegex.test(field.value)) {
                this.showError(field, 'Please enter a valid email address');
                return false;
            }
        } else if (field.required && field.value.trim() === '') {
            this.showError(field);
            return false;
        }

        this.hideError(field);
        field.classList.add('is-valid');
        return true;
    }

    validatePassword() {
        const password = this.passwordInput.value;
        const isStrong = password.length >= 8 &&
            /[A-Z]/.test(password) &&
            /[a-z]/.test(password) &&
            /\d/.test(password) &&
            /[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/.test(password);

        if (!isStrong) {
            this.showError(this.passwordInput, 'Password must meet all requirements');
            return false;
        }

        this.hideError(this.passwordInput);
        this.passwordInput.classList.add('is-valid');
        return true;
    }

    handleSubmit(e) {
        e.preventDefault();

        let isValid = true;
        const inputs = this.form.querySelectorAll('input[required]');

        // Validate all fields
        inputs.forEach(input => {
            if (!this.validateField(input)) {
                isValid = false;
            }
        });

        // Special validation for password
        if (!this.validatePassword()) {
            isValid = false;
        }

        // Check if passwords match
        if (this.passwordInput.value !== this.confirmPasswordInput.value) {
            this.showError(this.confirmPasswordInput, 'Passwords do not match');
            isValid = false;
        }

        if (isValid) {
            // Show success message (in a real app, this would submit to server)
            alert('Account created successfully! Welcome to TasklyGo!');

            // Reset form
            this.form.reset();
            this.form.querySelectorAll('.is-valid, .is-invalid').forEach(el => {
                el.classList.remove('is-valid', 'is-invalid');
            });

            // Hide all error messages
            this.form.querySelectorAll('.invalid-feedback').forEach(el => {
                el.classList.remove('show');
            });

            // Reset password strength indicators
            this.strengthContainer.classList.remove('show');
            this.requirementsList.classList.remove('show');
            this.strengthBar.style.width = '0%';
            this.strengthBar.className = 'progress-bar';
            this.strengthText.textContent = 'Enter password to check strength';
            this.matchIndicator.innerHTML = '';

            // Reset requirements
            Object.values(this.requirements).forEach(req => {
                req.classList.remove('met');
                req.classList.add('unmet');
                req.querySelector('i').className = 'fas fa-times';
            });
        }
    }
}

// Initialize the password validator when the page loads
document.addEventListener('DOMContentLoaded', () => {
    new PasswordValidator();
});