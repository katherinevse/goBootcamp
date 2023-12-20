class lengthOfLastWord {
    public int length_last(String s) {
        String str = s.trim(); //удаление пробелов
        int count = 0;
        for(int i=str.length()-1;i>=0;i--){
            if(str.charAt(i) != ' '){ //посимвольно считывает 
                count++;
            }
            else{
                break; 
            }
        }
        return count;
    }
}