let jsonData = `[
    {
        "id": 1,
        "name": "Laptop",
        "category": "Electronics",
        "price": 1000,
        "available": true
    },
    {
        "id": 2,
        "name": "Smartphone",
        "category": "Electronics",
        "price": 700,
        "available": false
    },
    {
        "id": 3,
        "name": "Induction Stove",
        "category": "Home Appliances",
        "price": 50,
        "available": true
    }
]`;
function parseJsonData(json) {
    return JSON.parse(json);
}

let products = parseJsonData(jsonData);

function addProduct(newProduct) {
    products.push(newProduct);
}

function updateProductPrice(productId, newPrice) {
    const product = products.find(p => p.id === productId);
    if (product) {
        product.price = newPrice;
    } else {
        return "Error: Product not found.";
    }
}

function getAvailableProducts() {
    return products.filter(p => p.available);
}

function getProductsByCategory(category) {
    return products.filter(p => p.category === category);
}

// Example Usage
// Adding a new product
addProduct({
    id: 4,
    name: "Blender",
    category: "Home Appliances",
    price: 29.99,
    available: true
});

// Updating the price of a product
console.log(updateProductPrice(2, 649.99)); // Updates price of Smartphone

// Getting available products
console.log(getAvailableProducts()); // Returns products that are available

// Getting products by category
console.log(getProductsByCategory("Electronics")); // Returns products in Electronics category

// Output the current products list
console.log(products);