class convertTemperature {
    public double[] converting(double celsius) {
        double Kelvin = celsius + 273.15;
        double Fahrenheit = celsius * 1.80 + 32.00;

        double ans[] = {Kelvin, Fahrenheit};

        return ans;

    }

    public static void main(String[] args) {
        convertTemperature converter = new convertTemperature();

        // Предположим, у нас есть температура в градусах Цельсия
        double celsiusTemperature = 25.0;

        // Используем метод converting для конвертации
        double[] result = converter.converting(celsiusTemperature);

        // Выводим результаты
        System.out.println("Температура в Кельвинах: " + result[0]);
        System.out.println("Температура в Фаренгейтах: " + result[1]);
    }
}