// Create 'customers' collection with validation schema
db.createCollection("customers", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["name", "email", "address", "phone", "registration_date"],
      properties: {
        name: { bsonType: "string", description: "must be a string and is required" },
        email: { bsonType: "string", pattern: "^.+@.+\..+$", description: "must be a valid email format and is required" },
        address: {
          bsonType: "object",
          required: ["street", "city", "zipcode"],
          properties: {
            street: { bsonType: "string", description: "must be a string and is required" },
            city: { bsonType: "string", description: "must be a string and is required" },
            zipcode: { bsonType: "string", pattern: "^[0-9]{5}$", description: "must be a 5-digit string and is required" }
          }
        },
        phone: { bsonType: "string", pattern: "^[0-9]{3}-[0-9]{4}$", description: "must be a valid phone number format (e.g., 555-1234)" },
        registration_date: { bsonType: "date", description: "must be a date and is required" }
      }
    }
  }
});

// Create 'orders' collection with validation schema
db.createCollection("orders", {
  validator: {
    $jsonSchema: {
      bsonType: "object",
      required: ["order_id", "customer_id", "order_date", "status", "items", "total_value"],
      properties: {
        order_id: { bsonType: "string", description: "must be a string and is required" },
        customer_id: { bsonType: "objectId", description: "must be an ObjectId referencing a customer document" },
        order_date: { bsonType: "date", description: "must be a date and is required" },
        status: { bsonType: "string", enum: ["pending", "shipped", "delivered", "canceled"], description: "must be one of the enum values and is required" },
        items: {
          bsonType: "array",
          minItems: 1,
          items: {
            bsonType: "object",
            required: ["product_name", "quantity", "price"],
            properties: {
              product_name: { bsonType: "string", description: "must be a string and is required" },
              quantity: { bsonType: "int", minimum: 1, description: "must be an integer greater than 0 and is required" },
              price: { bsonType: "double", minimum: 0, description: "must be a positive double and is required" }
            }
          }
        },
        total_value: { bsonType: "double", description: "must be a double and is required" }
      }
    }
  }
});

// Insert 5 customer documents into 'customers' collection
const customerDocs = [
  { name: "John Doe", email: "johndoe@example.com", address: { street: "123 Main St", city: "Springfield", zipcode: "12345" }, phone: "555-1234", registration_date: new Date("2023-01-01T12:00:00Z") },
  { name: "Jane Smith", email: "janesmith@example.com", address: { street: "456 Oak Ave", city: "Greenville", zipcode: "67890" }, phone: "555-5678", registration_date: new Date("2023-02-15T15:30:00Z") },
  { name: "Alice Brown", email: "alicebrown@example.com", address: { street: "789 Pine Rd", city: "Lakeside", zipcode: "11223" }, phone: "555-9876", registration_date: new Date("2023-03-12T09:45:00Z") },
  { name: "Bob White", email: "bobwhite@example.com", address: { street: "101 Maple Ln", city: "Riverside", zipcode: "44567" }, phone: "555-2345", registration_date: new Date("2023-04-20T14:20:00Z") },
  { name: "Charlie Green", email: "charliegreen@example.com", address: { street: "202 Birch Blvd", city: "Mountainview", zipcode: "55678" }, phone: "555-3456", registration_date: new Date("2023-05-05T10:00:00Z") }
];
db.customers.insertMany(customerDocs);

// Insert 5 order documents into 'orders' collection, each linked to a customer using the customer_id field
const ordersDocs = [
  { order_id: "ORD100001", customer_id: db.customers.findOne({ name: "John Doe" })._id, order_date: new Date("2023-10-15T12:00:00Z"), status: "shipped", items: [{ product_name: "Laptop", quantity: 1, price: 1500 }, { product_name: "Mouse", quantity: 2, price: 25 }], total_value: 1550 },
  { order_id: "ORD100002", customer_id: db.customers.findOne({ name: "Jane Smith" })._id, order_date: new Date("2023-11-01T12:00:00Z"), status: "shipped", items: [{ product_name: "Smartphone", quantity: 1, price: 600 }, { product_name: "Headphones", quantity: 1, price: 100 }], total_value: 700 },
  { order_id: "ORD100003", customer_id: db.customers.findOne({ name: "Alice Brown" })._id, order_date: new Date("2023-11-05T12:00:00Z"), status: "pending", items: [{ product_name: "Tablet", quantity: 1, price: 400 }], total_value: 400 },
  { order_id: "ORD100004", customer_id: db.customers.findOne({ name: "Bob White" })._id, order_date: new Date("2023-11-10T12:00:00Z"), status: "delivered", items: [{ product_name: "Monitor", quantity: 2, price: 300 }], total_value: 600 },
  { order_id: "ORD100005", customer_id: db.customers.findOne({ name: "Charlie Green" })._id, order_date: new Date("2023-11-12T12:00:00Z"), status: "shipped", items: [{ product_name: "Smartwatch", quantity: 1, price: 150 }], total_value: 150 }
];
db.orders.insertMany(ordersDocs);

// 2. Find Orders for a Specific Customer
const johnDoe = db.customers.findOne({ name: "John Doe" });
const johnOrders = db.orders.find({ customer_id: johnDoe._id }).toArray();
print("Orders for John Doe:");
printjson(johnOrders);

// 3. Find the Customer for a Specific Order
const order = db.orders.findOne({ order_id: "ORD100001" });
const customerForOrder = db.customers.findOne({ _id: order.customer_id });
print("Customer for Order ORD100001:");
printjson(customerForOrder);

// 4. Update Order Status
db.orders.updateOne({ order_id: "ORD100003" }, { $set: { status: "delivered" } });
print("Order status updated to 'delivered' for ORD100003.");

// 5. Delete an Order
db.orders.deleteOne({ order_id: "ORD100005" });
print("Order ORD100005 deleted.");
