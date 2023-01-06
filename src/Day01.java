/**
 * Created by 王盈 on 2023/1/6.
 * 给你一个正整数 num ，请你统计并返回 小于或等于 num 且各位数字之和为 偶数 的正整数的数目。
 * 正整数的 各位数字之和 是其所有位上的对应数字相加的结果。
 */
class Day01 {
    public int countEven(int num) {
        int sum = 0;
        if (num >= 1 && num < 10 ) {
            return lessTen(num);
        } else if (num >= 10 && num < 100) {
            return lessHundred(num);
        } else if (num >= 100 && num <1000) {
            return lessThousand(num);
        }else if(num==1000){
            return lessHundred(999);
        }
        return sum;
    }

    public int lessTen(int num) {
        int sum = 0;
        for (int i = 1; i <= num; i++) {
            if (i % 2 == 0) {
                sum++;
            }
        }
        System.out.println(sum);
        return sum;
    }
    public int lessHundred(int num){
        int sum = 0;
        for (int i = 1; i <= num; i++) {
            if(i >= 1 && i < 10 ){
                if (i % 2 == 0) {
                    sum++;
                }
            }else {
                //个位数
                int a = i % 10;
                //十位数
                int b = i / 10;
                int c = a + b;
                if (c % 2 == 0) {
                    sum++;
                }
            }
        }
        System.out.println(sum);
        return sum;
    }
    public int lessThousand(int num){
        int sum = 0;
        for (int i = 1; i <= num; i++) {
            if(i >= 1 && i < 10 ){
                if (i % 2 == 0) {
                    sum++;
                }
            }else if(i >= 10 && i < 100){
                //个位数
                int a = i % 10;
                //十位数
                int b = i / 10;
                int c = a + b;
                if (c % 2 == 0) {
                    sum++;
                }
            }else {
                int a1 = (i % 100) % 10;
                int b1 = (i % 100) / 10;
                int c1 = (i / 100);
                int d = a1 + b1 + c1;
                if (d % 2 == 0) {
                    sum++;
                }
            }
        }
        System.out.println(sum);
        return sum;
    }

    public static void main(String[] args) {
        Day01 d=new Day01();
        d.countEven(1000);
    }
}
