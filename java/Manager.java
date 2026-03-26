package java;

public class Manager extends Employee implements Workable{
    public Manager(String name){
        super(name);
    }
    @Override
    public void work(){
        System.out.println(name + "управляет");
    }
}
