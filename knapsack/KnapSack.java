import java.util.Arrays;
import java.util.Comparator;

class Item {
    int profit, weight;
    Item(int weight, int profit) {
        this.weight = weight;
        this.profit = profit;
    }
}

public class KnapSack {
    static double maxProfit(int W, Item arr[], int n) {
        // Sort items in by profit/weight ratio in descending order
        Arrays.sort(arr, new Comparator<Item>() {
            @Override
            public int compare(Item o1, Item o2) {
                return Double.compare(o2.profit / (double) o2.weight, o1.profit / (double) o1.weight);
            }
        });

        int currentWeight = 0;  // Current weight in knapsack
        double finalProfit = 0.0;  // Result (value in Knapsack)

        // Loop over items
        for (int i = 0; i < n; i++) {
            // If adding Item i will not out weight, add it completely
            if (currentWeight + arr[i].weight <= W) {
                currentWeight += arr[i].weight;
                finalProfit += arr[i].profit;
            } 
            else {
                // If we can't add entire item, add fractional part of it
                int remain = W - currentWeight;
                finalProfit += arr[i].profit * ((double) remain / arr[i].weight);
                break;  // As this is last item
            }
        }
      
        return finalProfit;
    }
  
    public static void main(String[] args) {
        int maxWeight = 50;  // Maximum weight the truck can carry
        Item arr[] = {new Item(10, 60), new Item(20, 100), new Item(30, 120)};
        int n = 3;  // Number of items
      
        System.out.println("Maximum profit we can obtain = " + 
                           maxProfit(maxWeight, arr, n));
    }
}