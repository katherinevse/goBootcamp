import java.util.Scanner;
import java.util.Arrays;

class  buyChoco{
    public int calculating(int[] prices, int money) {
        Arrays.sort(prices);
        System.out.println("Отсортированный массив: " + Arrays.toString(prices));

        int value = prices[0] + prices[1];
        if (money >= value) {
            return money-value;
        } else {
            return money;
        }

    }

    public static void main(String[] args){
        Scanner iScan = new Scanner(System.in);
        System.out.println("Введите номинал: " );

        int money = iScan.nextInt();
        int[] prices = new int[4];
        System.out.println("Введите массив: " );

        for (int i = 0; i < 4; i++) {
            prices[i] = iScan.nextInt();
        }
        buyChoco chocoCalculator = new buyChoco();
        int sortedPrices = chocoCalculator.calculating(prices, money);

    }
}

