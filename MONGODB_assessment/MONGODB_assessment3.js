// 1. Find All Customers Who Placed Orders in the Last Month
db.customers.aggregate([
  {
    $lookup: {
      from: "orders",
      localField: "_id",
      foreignField: "customer_id",
      as: "orders"
    }
  },
  {
    $unwind: "$orders"
  },
  {
    $match: {
      "orders.order_date": { 
        $gte: new Date(new Date().setDate(new Date().getDate() - 30)) // Last 30 days
      }
    }
  },
  {
    $sort: { "orders.order_date": -1 }
  },
  {
    $group: {
      _id: "$_id",
      name: { $first: "$name" },
      email: { $first: "$email" },
      most_recent_order_date: { $first: "$orders.order_date" }
    }
  },
  {
    $project: {
      name: 1,
      email: 1,
      most_recent_order_date: 1
    }
  }
]).forEach(function(customer) {
  print("Customer: " + customer.name + ", Email: " + customer.email + ", Most Recent Order Date: " + customer.most_recent_order_date);
});

// 2. Find All Products Ordered by a Specific Customer (John Doe)
db.orders.aggregate([
  {
    $lookup: {
      from: "customers",
      localField: "customer_id",
      foreignField: "_id",
      as: "customer"
    }
  },
  {
    $unwind: "$customer"
  },
  {
    $match: {
      "customer.name": "John Doe"
    }
  },
  {
    $unwind: "$items"
  },
  {
    $group: {
      _id: "$items.product_name",
      total_quantity: { $sum: "$items.quantity" }
    }
  },
  {
    $project: {
      product_name: "$_id",
      total_quantity: 1
    }
  }
]).forEach(function(product) {
  print("Product: " + product.product_name + ", Total Quantity Ordered: " + product.total_quantity);
});

// 3. Find the Top 3 Customers with the Most Expensive Total Orders
db.orders.aggregate([
  {
    $lookup: {
      from: "customers",
      localField: "customer_id",
      foreignField: "_id",
      as: "customer"
    }
  },
  {
    $unwind: "$customer"
  },
  {
    $group: {
      _id: "$customer._id",
      name: { $first: "$customer.name" },
      total_order_value: { $sum: "$total_value" }
    }
  },
  {
    $sort: { total_order_value: -1 }
  },
  {
    $limit: 3
  },
  {
    $project: {
      name: 1,
      total_order_value: 1
    }
  }
]).forEach(function(customer) {
  print("Customer: " + customer.name + ", Total Order Value: $" + customer.total_order_value);
});

// 4. Add a New Order for an Existing Customer (Jane Smith)
var janeSmith = db.customers.findOne({ name: "Jane Smith" });

if (janeSmith) {
  var newOrder = {
    customer_id: janeSmith._id,
    order_date: new Date(),
    status: "processing",
    total_value: 1200,
    items: [
      { product_name: "Smartphone", quantity: 1, price: 800 },
      { product_name: "Headphones", quantity: 1, price: 400 }
    ]
  };

  db.orders.insertOne(newOrder);
  print("New order added for Jane Smith with items: Smartphone and Headphones.");
} else {
  print("Customer Jane Smith not found.");
}
