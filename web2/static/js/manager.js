// Set default due date to tomorrow
const tomorrow = new Date();
tomorrow.setDate(tomorrow.getDate() + 1);
document.getElementById('due_date').valueAsDate = tomorrow;

// Simple filter highlighting based on URL params
const urlParams = new URLSearchParams(window.location.search);
const filter = urlParams.get('filter');
if (filter) {
    document.querySelectorAll('.btn-group a').forEach(btn => {
        btn.classList.remove('active');
        if (btn.href.includes(`filter=${filter}`)) {
            btn.classList.add('active');
        }
    });
}