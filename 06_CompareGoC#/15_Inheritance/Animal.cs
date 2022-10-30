using System;

namespace Inheritance {

    /******************************************************************************
    ***************************** Grandparent class *******************************
    ******************************************************************************/
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

        public void Eat() {
            Console.WriteLine($"Eat: ");
        }

        public void Move() {
            Console.WriteLine($"Move: ");
        }

        // If you want define a method that can create a object 
        // and child class can override this method
        // ===> use virtual
        public virtual void MakeNoise() { 
            Console.WriteLine($"Make some noise: ");
        }

        // If you want define a method that don't create a object 
        // and all child classes must override this method
        // ===> use abstract
        // public abstract void Action();

        // Method to use attribute
        public void ShowInfo()
        {
            Console.WriteLine($"This aminal: Gender = {gender}; Age = {age}; Weigth = {weigth} ");
        }

        // Implement method overloading
        public void ShowInfo(string act)
        {
            Console.WriteLine($"This aminal: Gender = {gender}; Age = {age}; Weigth = {weigth} ");
            
            if(String.Equals(act, "sleep")) {
                Sleep();
            } else if(String.Equals(act, "eat")) {
                Eat();
            } else if(String.Equals(act, "move")) {
                Move();  
            } else if(String.Equals(act, "all")) {
                Sleep();
                Eat();
                Move();
            } else {
                Console.WriteLine($"Please choose action: sleep, eat, move, all");
            }
        }
    } 


    /************************** Inheritance Mammal: Parent class *******************************/
    public class Mammal : Animals {
        protected int Legs {get; set;} // Add a new attribute

        // Avoid error 1: The type `Inheritance.Mammal' does not contain a constructor that takes `0' arguments
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

    /************* Cat class: Child class ************/
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

    /************************** Inheritance Bird: Parent class *******************************/

    public class Bird : Animals {
        public string color; // thuộc tính mới thêm
        public Bird() {
            color = "Yellow";
        }

        // Override virtual method 
        public override void MakeNoise() {
            // Before override method, we can call virtual method behavior
            // Call virtual from base class to get some nose.
            base.MakeNoise();

            // Redefine method 
            Console.WriteLine($"Bird sing ");
        }
    }

    /************************** Inheritance Reptile: Parent class *******************************/
    public class Reptile : Animals {
 
    }


    /************************** Inheritance Fish: Parent class *******************************/
    public class Fish : Animals {

    }

    /*****************************************************************************************
    ***************************** Strategy Design Pattern  ***********************************
    ******************************************************************************************/ 
    /************************** Interface and Concrete *******************************/
    /*** Fly Behavior ***/
    public interface FlyBehavior {
        void fly();
    }

    class FlyNoWay : FlyBehavior {
        public void fly() {
            Console.WriteLine($"I can't fly");
        }
    }

    class FlyWithWings : FlyBehavior {
        public void fly() {
            Console.WriteLine("I'm flying!!");
        }
    }

    /*** Quack Behavior ***/
    public interface QuackBehavior {
        void quack();
    }

    class Quack : QuackBehavior {
        public void quack() {
            Console.WriteLine($"Quack");
        }
    }

    /************************** Inheritance Duck: Parent class *******************************/
    public abstract class Duck : Animals {
        protected FlyBehavior flyBehavior;
	    protected QuackBehavior quackBehavior;

        public Duck() {
        }

        public void setFlyBehavior(FlyBehavior fb) {
            flyBehavior = fb;
        }

        public void setQuackBehavior(QuackBehavior qb) {
            quackBehavior = qb;
        }

        public abstract void display();

        public void performFly() {
            flyBehavior.fly();
        }

        public void performQuack() {
            quackBehavior.quack();
        }

        public void swim() {
            Console.WriteLine($"All ducks float, even decoys!");
        }

    }

    /************************** Inheritance Duck: Child class *******************************/
    public class MallardDuck : Duck {

        public MallardDuck() {

            base.quackBehavior = new Quack();
            base.flyBehavior = new FlyWithWings();

        }

        public override void display() {
            Console.WriteLine($"I'm a real Mallard duck");
        }
    }


    /******************************************************************************
    ********************************** Program ************************************
    ******************************************************************************/
    class Program {
        static void Main (string[] args) {

            Console.WriteLine($"*********** Example 1: Inheritance ***********");
            /* Example 1: Inheritance */  
              
            Cat cat = new Cat();
            cat.ShowInfo(); // Phương thức này kế thừa từ lớp cơ sở
            cat.ShowLegs();
            cat.SetLegs(3);
            cat.ShowLegs();
            cat.Eat(); // phương thức của riêng Cat
            cat.ShowInfo("sleep");


            Console.WriteLine($"*********** Example 2: Override method ***********");
            /* Example 2: Override method */    
            Bird bird = new Bird();
            bird.MakeNoise();

            Console.WriteLine($"*********** Example 3: Abstract method ***********");
            /* Example 3: Abstract method */

            Console.WriteLine($"*********** Example 4: Stategy Design Pattern ***********");
            /* Example 4: Stategy Design Pattern */
            MallardDuck	mallard = new MallardDuck();
            mallard.display();
            mallard.performFly();
            mallard.performQuack();
        }       

    }
}