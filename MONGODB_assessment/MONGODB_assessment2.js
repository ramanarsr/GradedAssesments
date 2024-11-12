// 1. Calculate Total Value of All Orders by Customer
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
    $project: {
      name: 1,
      total_value: { 
        $sum: "$orders.total_value" 
      }
    }
  },
  {
    $sort: { total_value: -1 }
  },
  {
    $project: {
      name: 1,
      total_value: 1
    }
  }
]).forEach(function(customer) {
  print("Customer: " + customer.name + ", Total Order Value: $" + customer.total_value);
});

// 2. Group Orders by Status
db.orders.aggregate([
  {
    $group: {
      _id: "$status",
      order_count: { $sum: 1 }
    }
  },
  {
    $sort: { order_count: -1 }
  }
]).forEach(function(statusGroup) {
  print("Status: " + statusGroup._id + ", Orders Count: " + statusGroup.order_count);
});

// 3. List Customers with Their Recent Orders
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
    $sort: { "orders.order_date": -1 }
  },
  {
    $group: {
      _id: "$_id",
      name: { $first: "$name" },
      email: { $first: "$email" },
      recent_order: { $first: "$orders" }
    }
  },
  {
    $project: {
      name: 1,
      email: 1,
      order_id: "$recent_order.order_id",
      total_value: "$recent_order.total_value"
    }
  }
]).forEach(function(customerOrder) {
  print("Customer: " + customerOrder.name + ", Email: " + customerOrder.email + ", Recent Order ID: " + customerOrder.order_id + ", Total Value: $" + customerOrder.total_value);
});

// 4. Find the Most Expensive Order by Customer
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
    $project: {
      name: 1,
      orders: 1
    }
  },
  {
    $addFields: {
      most_expensive_order: {
        $arrayElemAt: [
          {
            $sortArray: {
              input: "$orders",
              sortBy: { total_value: -1 }
            }
          },
          0
        ]
      }
    }
  },
  {
    $project: {
      name: 1,
      most_expensive_order_id: "$most_expensive_order.order_id",
      most_expensive_order_value: "$most_expensive_order.total_value"
    }
  }
]).forEach(function(expensiveOrder) {
  print("Customer: " + expensiveOrder.name + ", Most Expensive Order ID: " + expensiveOrder.most_expensive_order_id + ", Total Value: $" + expensiveOrder.most_expensive_order_value);
});
