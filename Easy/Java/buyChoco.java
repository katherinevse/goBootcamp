import java.util.Scanner;
import java.util.Arrays;

class  buyChoco{
    public int[] calculating(int[] prices) {
        Arrays.sort(prices);
        System.out.println("Отсортированный массив: " + Arrays.toString(prices));
        return prices;
    }

    public static void main(String[] args){
        Scanner iScan = new Scanner(System.in);
//        int money = iScan.nextInt();
        int[] prices = new int[4];
        System.out.println("Введите массив: " );

        for (int i = 0; i < 4; i++) {
            prices[i] = iScan.nextInt();
        }
        buyChoco chocoCalculator = new buyChoco();
        int[] sortedPrices = chocoCalculator.calculating(prices);

    }
}

