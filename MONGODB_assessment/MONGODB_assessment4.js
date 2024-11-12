// 1. Find Customers Who Have Not Placed Orders
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
    $match: {
      "orders": { $size: 0 } // Customers who have no orders
    }
  },
  {
    $project: {
      name: 1,
      email: 1
    }
  }
]).forEach(function(customer) {
  print("Customer Name: " + customer.name + ", Email: " + customer.email);
});

// 2. Calculate the Average Number of Items Ordered per Order
db.orders.aggregate([
  {
    $unwind: "$items"
  },
  {
    $group: {
      _id: "$_id", // Grouping by order ID to count total items per order
      total_items: { $sum: "$items.quantity" }
    }
  },
  {
    $group: {
      _id: null, // Grouping all orders together to calculate the average
      average_items: { $avg: "$total_items" }
    }
  },
  {
    $project: {
      average_items: 1
    }
  }
]).forEach(function(result) {
  print("Average Number of Items Ordered per Order: " + result.average_items);
});

// 3. Join Customer and Order Data Using $lookup
db.customers.aggregate([
  {
    $lookup: {
      from: "orders",
      localField: "_id",
      foreignField: "customer_id",
      as: "order_details"
    }
  },
  {
    $unwind: "$order_details" // Unwind orders to show each order separately
  },
  {
    $project: {
      name: 1,
      email: 1,
      "order_details.order_id": 1,
      "order_details.total_value": 1,
      "order_details.order_date": 1
    }
  }
]).forEach(function(customerOrder) {
  print("Customer Name: " + customerOrder.name + ", Email: " + customerOrder.email +
        ", Order ID: " + customerOrder.order_details.order_id + ", Total Value: " + customerOrder.order_details.total_value + 
        ", Order Date: " + customerOrder.order_details.order_date);
});
