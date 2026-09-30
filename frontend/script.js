document.addEventListener('DOMContentLoaded', () => {
    const API_BASE = 'http://localhost:8080';
    
    let currentSection = 'collections';
    
    const sections = {
        collections: document.getElementById('collections'),
        customers: document.getElementById('customers'),
        measurements: document.getElementById('measurements'),
        orders: document.getElementById('orders'),
        tailors: document.getElementById('tailors')
    };

    const customerModal = document.getElementById('addCustomerModal');
    const tailorModal = document.getElementById('addTailorModal');

    // Toggle form functions
    window.showAddCustomerForm = () => {
        customerModal.style.display = 'flex';
    };

    window.hideAddCustomerForm = () => {
        customerModal.style.display = 'none';
    };

    window.showAddTailorForm = () => {
        tailorModal.style.display = 'flex';
    };

    window.hideAddTailorForm = () => {
        tailorModal.style.display = 'none';
    };

    // Close modals on outside click
    customerModal.addEventListener('click', (e) => {
        if (e.target === customerModal) {
            hideAddCustomerForm();
        }
    });

    tailorModal.addEventListener('click', (e) => {
        if (e.target === tailorModal) {
            hideAddTailorForm();
        }
    });

    // Switch between sections
    window.switchSection = (sectionId) => {
        // Hide all sections
        Object.values(sections).forEach(section => {
            section.style.display = 'none';
            section.classList.remove('active');
        });
        
        // Show selected section
        if (sections[sectionId]) {
            sections[sectionId].style.display = 'block';
            sections[sectionId].classList.add('active');
        }
        
        currentSection = sectionId;
        updateNavLink(sectionId);
    };

    function updateNavLink(sectionId) {
        const links = document.querySelectorAll('.nav-links a');
        links.forEach(link => {
            link.classList.remove('active');
            if (link.getAttribute('href') === `#${sectionId}`) {
                link.classList.add('active');
            }
        });
    }

    // Nav link clicks
    document.querySelectorAll('.nav-links a').forEach(link => {
        link.addEventListener('click', (e) => {
            e.preventDefault();
            const sectionId = link.getAttribute('href').replace('#', '');
            switchSection(sectionId);
        });
    });

    // Initial load - show collections
    let customers = [];

    async function loadCustomers() {
        try {
            const response = await fetch(`${API_BASE}/api/customers`);
            customers = await response.json();
            renderCustomers();
        } catch (error) {
            console.error('Error loading customers:', error);
            document.getElementById('customersGrid').innerHTML = '<div class="error">Error loading customers. Make sure the Go backend is running.</div>';
        }
    }

    function renderCustomers() {
        const grid = document.getElementById('customersGrid');
        if (customers.length === 0) {
            grid.innerHTML = '<div class="empty">No customers found. Click "+ Add Customer" to add one.</div>';
            return;
        }
        
        grid.innerHTML = customers.map(customer => `
            <div class="customer-card">
                <div class="card-header">Customer: ${escapeHtml(customer.name)}</div>
                <div class="card-meta">
                    <span>Email: ${escapeHtml(customer.email)}</span>
                    <span>Phone: ${escapeHtml(customer.phone)}</span>
                </div>
            </div>
        `).join('');
    }

    // Add customer form submission
    document.getElementById('customerForm').addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const name = document.getElementById('customerName').value;
        const email = document.getElementById('customerEmail').value;
        const phone = document.getElementById('customerPhone').value;
        
        try {
            await fetch(`${API_BASE}/api/customers`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ name, email, phone })
            });
            
            hideAddCustomerForm();
            document.getElementById('customerName').value = '';
            document.getElementById('customerEmail').value = '';
            document.getElementById('customerPhone').value = '';
            
            loadCustomers();
        } catch (error) {
            console.error('Error adding customer:', error);
            alert('Error adding customer. Make sure the Go backend is running.');
        }
    });

    // Load measurements
    async function loadMeasurements() {
        try {
            const response = await fetch(`${API_BASE}/api/measurement-profiles`);
            const measurements = await response.json();
            renderMeasurements(measurements);
        } catch (error) {
            console.error('Error loading measurements:', error);
            document.getElementById('measurementsGrid').innerHTML = '<div class="error">Error loading measurements.</div>';
        }
    }

    function renderMeasurements(measurements) {
        const grid = document.getElementById('measurementsGrid');
        if (measurements.length === 0) {
            grid.innerHTML = '<div class="empty">No measurement profiles found.</div>';
            return;
        }
        
        grid.innerHTML = measurements.map(m => `
            <div class="measurement-card">
                <div class="card-header">Measurement: ${escapeHtml(m.label)}</div>
                <div class="card-meta">
                    <span>Customer ID: ${escapeHtml(m.customer_id)}</span>
                    <span>Bust: ${m.bust}cm</span>
                    <span>Waist: ${m.waist}cm</span>
                    <span>Hips: ${m.hips}cm</span>
                </div>
            </div>
        `).join('');
    }

    // Load orders
    async function loadOrders() {
        try {
            const response = await fetch(`${API_BASE}/api/orders`);
            const orders = await response.json();
            renderOrders(orders);
        } catch (error) {
            console.error('Error loading orders:', error);
            document.getElementById('ordersGrid').innerHTML = '<div class="error">Error loading orders.</div>';
        }
    }

    function renderOrders(orders) {
        const grid = document.getElementById('ordersGrid');
        if (orders.length === 0) {
            grid.innerHTML = '<div class="empty">No orders found.</div>';
            return;
        }
        
        grid.innerHTML = orders.map(order => `
            <div class="order-card">
                <div class="card-header">Order #${order.id.substring(0, 8)}...</div>
                <div class="card-meta">
                    <span>Customer: ${order.customer_id.substring(0, 8)}...</span>
                    <span>Tailor: ${order.tailor_id.substring(0, 8)}...</span>
                    <span>Status: ${order.status}</span>
                    <span>Price: $${order.price || 0}</span>
                </div>
            </div>
        `).join('');
    }

    // Load tailers
    async function loadTailors() {
        try {
            const response = await fetch(`${API_BASE}/api/tailors`);
            const tailors = await response.json();
            renderTailors(tailors);
        } catch (error) {
            console.error('Error loading tailors:', error);
            document.getElementById('tailorsGrid').innerHTML = '<div class="error">Error loading tailors.</div>';
        }
    }

    function renderTailors(tailors) {
        const grid = document.getElementById('tailorsGrid');
        if (tailors.length === 0) {
            grid.innerHTML = '<div class="empty">No tailors found. Click "+ Add Tailor" to add one.</div>';
            return;
        }
        
        grid.innerHTML = tailors.map(tailor => `
            <div class="tailor-card">
                <div class="card-header">Tailor: ${escapeHtml(tailor.name)}</div>
                <div class="card-meta">
                    <span>Phone: ${escapeHtml(tailor.phone)}</span>
                    <span>Email: ${escapeHtml(tailor.email)}</span>
                </div>
            </div>
        `).join('');
    }

    // Initialize - load data when page loads
    async function init() {
        // Show collections by default
        switchSection('collections');
        
        // Load all data
        await loadCustomers();
        await loadMeasurements();
        await loadOrders();
        await loadTailors();
    }

    function escapeHtml(text) {
        const div = document.createElement('div');
        div.textContent = text;
        return div.innerHTML;
    }

    init();
});