using System;

namespace Inheritance {

    public class Animals {
        private string gender {get; set;}
        private int age {get; set;}
        private int weigth {get; set;}

        // Constructor default with none parameter
        public Animals() { 

        }
        
        // Constructor parameter
        public Animals(string gender, int age, int weigth) {
            this.gender = gender;
            this.age = age;
            this.weigth = weigth;
        }

        //  Methods to inherit
        public void Sleep() {
            Console.WriteLine($"Sleep:");
        }
        public void Eat()
        {
            Console.WriteLine($"Eat: ");
        }
        public void Move()
        {
            Console.WriteLine($"Move: ");
        }

        // Method to use attribute
        public void ShowAge()
        {
            Console.WriteLine($"Age: {age}");
        }
    } 

    /************* Inheritance Mammal************/
    public class Mammal : Animals {
        public int Legs {get; set;} // Add a new attribute

        // Avoid error: The type `Inheritance.Mammal' does not contain a constructor that takes `0' arguments
        // Must implement this constructor
        public Mammal() {

        }
        public Mammal(string gender, int age, int weigth) : base(gender, age, weigth) {
            Legs = 4;
        }
        public void ShowLegs() {
            Console.WriteLine($"Legs: {Legs}");
        }
        public void SetLegs(int Legs) {
            this.Legs = Legs;
        }
    }

    public class Cat : Mammal {
        public string food;
        public Cat() {
            Legs = 4;
            food = "Mouse";
        }

        public Cat(string gender, int age, int weigth) : base(gender, age, weigth) {
             Legs = 4;
             food = "Mouse";
        }
    }

    /************* Inheritance Bird ************/
    public class Bird : Animals {
        public string color; // thuộc tính mới thêm
        public Bird() {
            color = "Yellow";
        }
    }

    /************* Inheritance Reptile************/
    public class Reptile : Animals {
 
    }


    /************* Inheritance Fish************/

    public class Fish : Animals {

    }

    class Program {
        static void Main (string[] args) {
            /* Ví dụ kế thừa */    
            Cat cat = new Cat();
            cat.ShowAge (); // Phương thức này kế thừa từ lớp cơ sở
            cat.SetLegs(3);
            cat.ShowLegs();
            cat.Eat (); // phương thức của riêng Cat

        }

    }
}