package java;

public class Developer extends Employee implements Workable{
    public Developer(String name){
        super(name);
    }
    @Override
    public void work(){
        System.out.println(name + "пишет код");
    }
}
