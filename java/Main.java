package java;

import java.util.ArrayList;

public class Main {
    public static void main(String[] args){
        ArrayList<Employee> employees = new ArrayList<>();

        employees.add(new Developer("Anvar"));
        employees.add(new Manager("Parananana"));

        for (Employee e : employees){
            System.out.println(e.name);
        }

        String search = "Anvar";
        for (Employee e : employees){
            if (e.name.equals(search)){
                System.out.println("Найден" + e.name);
            }
        }

        for (Employee e : employees){
            ((Workable) e).work();
        }
    }
}