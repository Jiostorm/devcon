
# **Memory Management Models**
## [_**Ownership**_](https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html)
<p><strong>Ownership</strong> is a set of rules that govern how a Rust program manages memory.
All programs have to manage the way they use a computer’s memory while running.
Some languages have garbage collection that regularly looks for no-longer-used memory as the program runs; in other languages,
the programmer must explicitly allocate and free the memory.
<br><br>
Rust uses a third approach: memory is managed through a system of ownership with a set of rules that the compiler checks.
If any of the rules are violated, the program won’t compile. None of the features of ownership will slow down your program while it’s running. <small><em>(from the Rust docs)</em></small>
</p>

### _**Ownership Rules**_
Each value in Rust has an owner.
There can only be one owner at a time.
When the owner goes out of scope, the value will be dropped.
- Each value in Rust has an owner.
- There can only be one owner at a time.
- When the owner goes out of scope, the value will be dropped.
## [_**Garbage Collection**_](https://en.wikipedia.org/wiki/Garbage_collection_(computer_science))
<p>
In Computer Science, <strong>Garbage Collection</strong> (GC) is a form of automatic memory management.
The garbage collector attempts to reclaim memory that was allocated by the program, but is no longer referenced; such memory is called garbage.
<br><br>
Garbage Collection relieves the programmer from doing manual memory management,
where the programmer specifies what objects to de-allocate and return to the memory system and when to do so. <small><em>(from the wikipedia)</em></small>
</p>

| Type | Pros | Cons |
| :--- | :--- | :-- |
| _`Ownership`_ | <pre>* Control over memory<br>* Free from potential bugs and errors<br>* Faster runtime</pre> | <pre>* Steep learning curve<br>* Slow coding time</pre> |
| _`Garbage Collection`_ | <pre>* Free from potential errors<br>* Faster coding time</pre> | <pre>* No control over memory allocation<br>* Unpredictable behavior since it is a non-deterministic mechanism</pre> |
| _`Manual Memory Management`_ | <pre>* Total control over memory allocation<br>* Faster runtime</pre> | <pre>* Prone to bugs and errors<br>* Developer-dependent performance</pre>|
