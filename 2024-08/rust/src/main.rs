// Ownership

#[derive(Debug, Clone)]
struct Person {
    name: String,
    age: u8,
}

impl Person {
    fn new(name: String, age: u8) -> Person {
        let person = Person { name, age };
        person
    }
}

pub fn main() {
    let jio = Person::new("Jio".to_string(), 23);
    let cla = Person::new("Cla".to_string(), 23);

    let persons: [Person; 2] = [jio.clone(), cla.clone()];
    println!("{:?}", persons);

    //owned_say_hello(jio);
    //owned_say_hello(cla);
    borrowed_say_hello(&jio);
    borrowed_say_hello(&cla);

    println!("{}, {}", jio.name, jio.age);
    println!("{}, {}", cla.name, cla.age);
}

#[allow(dead_code)]
fn owned_say_hello(person: Person) {
    println!("Hello {}!", person.name)
}

#[allow(dead_code)]
fn borrowed_say_hello(person: &Person) {
    println!("Hello {}!", person.name)
}
