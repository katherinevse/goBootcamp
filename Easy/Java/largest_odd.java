//substring - это метод в Java, который используется для извлечения подстроки из строки

class Solution {
    public String largestOddNumber(String num) {
        for (int i = num.length()-1; i >= 0; i--) {
            int current=num.charAt(i);
            if (current%2 !=0) {
                return num.substring(0,i+1);
                
            }
        }
        return "";
        
    }
}