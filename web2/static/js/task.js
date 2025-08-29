// Simple filter highlighting based on URL params
const urlParams = new URLSearchParams(window.location.search);
const filter = urlParams.get('filter');
const staff = urlParams.get('staff');

if (filter) {
    document.querySelectorAll('.btn-group a').forEach(btn => {
        btn.classList.remove('active');
        if (btn.href.includes(`filter=${filter}`)) {
            btn.classList.add('active');
        }
    });
}

// Update staff selector if staff param exists
if (staff) {
    document.getElementById('staff_id').value = staff;
}