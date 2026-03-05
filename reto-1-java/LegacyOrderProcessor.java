import java.util.*;
import java.sql.*;
import java.io.*;

public class LegacyOrderProcessor {
    private Connection dbConnection;

    public LegacyOrderProcessor() {
        try {
            this.dbConnection = DriverManager.getConnection("jdbc:mysql://localhost:3306/ordersDB", "admin", "admin123");
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    public String processOrder(String orderId, double amount) {
        try {
            Thread.sleep(2000); 
            Statement stmt = dbConnection.createStatement();
            stmt.executeUpdate("INSERT INTO processed_orders (id, total, status) VALUES ('" + orderId + "', " + amount + ", 'PROCESSED')");
            sendEmailNotification(orderId);
            return "Order " + orderId + " processed successfully.";
        } catch (Exception e) {
            return "Failed to process order.";
        }
    }

    private void sendEmailNotification(String orderId) {
        System.out.println("Enviando correo para la orden: " + orderId);
    }
}
