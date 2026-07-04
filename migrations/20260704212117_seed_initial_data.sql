-- +goose Up

-- Курсы
INSERT INTO courses (name, description) VALUES
                                            ('Golang Developer', 'A backend development course focused on building reliable services in Go, covering language fundamentals, control flow, and core data structures through practical exercises.'),
                                            ('Python Developer', 'An introductory course to Python programming, covering syntax fundamentals, core data structures, and writing reusable functions and modules.'),
                                            ('JavaScript Developer', 'A frontend-oriented course covering JavaScript fundamentals and browser interaction through the Document Object Model.');

-- Главы курса "Golang Developer" (course_id = 1)
INSERT INTO chapters (name, description, "order", course_id) VALUES
                                                                 ('Control Structures', 'This chapter introduces the constructs used to control the flow of execution in a Go program, including conditional branching and multi-way selection.', 1, 1),
                                                                 ('Data Types and Variables', 'This chapter covers the built-in data types available in Go and how variables and constants are declared and used.', 2, 1);

-- Главы курса "Python Developer" (course_id = 2)
INSERT INTO chapters (name, description, "order", course_id) VALUES
                                                                 ('Introduction to Python', 'This chapter introduces the basic syntax of Python and the core built-in data structures used to store collections of values.', 1, 2),
                                                                 ('Functions and Modules', 'This chapter explains how to define reusable blocks of code as functions and how to organize code across multiple files using modules.', 2, 2);

-- Главы курса "JavaScript Developer" (course_id = 3)
INSERT INTO chapters (name, description, "order", course_id) VALUES
                                                                 ('JavaScript Fundamentals', 'This chapter covers the basic building blocks of JavaScript, including variables, data types, functions, and scope.', 1, 3),
                                                                 ('DOM Manipulation', 'This chapter explains how JavaScript interacts with the Document Object Model to read and modify content on a web page.', 2, 3);

-- Уроки главы "Control Structures" (chapter_id = 1)
INSERT INTO lessons (name, description, content, "order", chapter_id) VALUES
                                                                          ('If-else Statement in Golang', 'Understanding conditional branching in Go.',
                                                                           'The if statement in Go evaluates a boolean condition and executes a block of code only if that condition is true. Unlike many other languages, Go does not require parentheses around the condition, but the opening brace must be on the same line as the condition. An if statement can be followed by one or more else if branches, and an optional final else branch that runs when none of the preceding conditions are true. Go also allows a short initialization statement before the condition, separated by a semicolon, which is commonly used to declare a variable scoped only to the if-else block. This keeps helper variables from leaking into the surrounding function scope.',
                                                                           1, 1),
                                                                          ('Switch Statement in Golang', 'Using switch as a cleaner alternative to multiple if-else chains.',
                                                                           'The switch statement in Go provides a way to compare a value against a list of possible matches without writing a long chain of if-else statements. Unlike in some other languages, Go automatically breaks out of a case after it executes, so the fallthrough keyword must be used explicitly if execution should continue into the next case. A switch statement can also be written without an expression, in which case each case contains its own boolean condition, making it a readable alternative to a long if-else chain. Switch statements can operate on many types, including integers, strings, and even types themselves through a type switch.',
                                                                           2, 1);

-- Уроки главы "Data Types and Variables" (chapter_id = 2)
INSERT INTO lessons (name, description, content, "order", chapter_id) VALUES
                                                                          ('Basic Data Types in Golang', 'Overview of Go''s built-in primitive types.',
                                                                           'Go is a statically typed language, meaning the type of every variable is known at compile time. The language provides several categories of built-in types: numeric types such as int, float64, and their fixed-size variants like int32 and int64; the bool type for true or false values; and the string type for representing text as an immutable sequence of bytes. Go also provides complex composite types built from these primitives, such as arrays, slices, and maps, which are covered in later chapters. Choosing the appropriate numeric type affects both memory usage and the range of values a variable can hold.',
                                                                           1, 2),
                                                                          ('Variables and Constants', 'Declaring and initializing variables and constants in Go.',
                                                                           'Variables in Go can be declared explicitly using the var keyword followed by a name and a type, or implicitly using the short declaration operator, which infers the type from the assigned value. Go requires that every declared variable be used somewhere in the code, otherwise the compiler raises an error. Constants are declared using the const keyword and must be assigned a value that can be determined at compile time, such as a literal number or string. Unlike variables, constants cannot be reassigned after declaration, which makes them useful for values that should remain fixed throughout the program, such as configuration limits or mathematical constants.',
                                                                           2, 2);

-- Уроки главы "Introduction to Python" (chapter_id = 3)
INSERT INTO lessons (name, description, content, "order", chapter_id) VALUES
                                                                          ('Python Syntax Basics', 'Core syntax rules that define how Python code is structured.',
                                                                           'Python uses indentation rather than curly braces to define blocks of code, which means consistent spacing is not just a style preference but a syntax requirement. A typical Python program consists of statements executed sequentially from top to bottom, with control structures such as if statements and loops introducing indented blocks. Python is dynamically typed, so a variable does not need an explicit type declaration; its type is determined by the value assigned to it at runtime. Comments in Python are written using the hash symbol, and multi-line strings enclosed in triple quotes are often used for longer documentation blocks.',
                                                                           1, 3),
                                                                          ('Working with Lists and Dictionaries', 'Two of the most commonly used built-in data structures in Python.',
                                                                           'A list in Python is an ordered, mutable collection that can hold items of different types, and elements can be added, removed, or modified after the list is created. Lists support indexing and slicing, allowing a program to access individual elements or a subrange of elements using square brackets. A dictionary stores data as key-value pairs, allowing fast lookup of a value given its associated key rather than a numeric position. Both structures support iteration using a for loop, and dictionaries additionally allow iterating over keys, values, or key-value pairs together.',
                                                                           2, 3);

-- Уроки главы "Functions and Modules" (chapter_id = 4)
INSERT INTO lessons (name, description, content, "order", chapter_id) VALUES
                                                                          ('Defining Functions in Python', 'How to write and call reusable blocks of code.',
                                                                           'A function in Python is defined using the def keyword, followed by a name, a parenthesized list of parameters, and a colon that begins an indented code block. Functions can accept default parameter values, which are used when the caller does not explicitly supply that argument. The return statement is used to send a value back to the caller; if a function has no explicit return statement, it implicitly returns the special value None. Python also supports variable-length argument lists using the asterisk syntax, allowing a function to accept an arbitrary number of positional or keyword arguments.',
                                                                           1, 4),
                                                                          ('Importing and Using Modules', 'Organizing and reusing code across multiple files.',
                                                                           'A module in Python is simply a file containing Python code that can be imported and reused in other files using the import statement. Python includes a large standard library of built-in modules covering tasks such as file handling, mathematics, and date and time manipulation. A specific function or object can be imported directly from a module using the from-import syntax, avoiding the need to reference the full module name each time it is used. Organizing related functions and classes into separate modules, and grouping modules into packages, helps keep larger projects maintainable as they grow.',
                                                                           2, 4);

-- Уроки главы "JavaScript Fundamentals" (chapter_id = 5)
INSERT INTO lessons (name, description, content, "order", chapter_id) VALUES
                                                                          ('Variables and Data Types in JavaScript', 'Declaring variables and understanding JavaScript''s core data types.',
                                                                           'JavaScript provides three keywords for declaring variables: var, let, and const. The let and const keywords were introduced to address scoping issues present in var, restricting a variable''s visibility to the block in which it is declared rather than the entire function. JavaScript is dynamically typed and includes primitive types such as number, string, boolean, undefined, and null, alongside the object type used for more complex structures. The const keyword declares a variable that cannot be reassigned after its initial value is set, though objects and arrays declared with const can still have their internal contents modified.',
                                                                           1, 5),
                                                                          ('Functions and Scope', 'How functions are declared and how scope determines variable visibility.',
                                                                           'Functions in JavaScript can be declared in several ways: as a function declaration, as a function expression assigned to a variable, or as an arrow function using a shorter syntax. Arrow functions differ from regular functions in how they handle the this keyword, inheriting it from the surrounding context rather than defining their own. Scope in JavaScript determines which parts of the code can access a given variable, with block scope applying to variables declared with let and const, and function scope applying to variables declared with var. Closures occur when a function retains access to variables from its containing scope even after that outer function has finished executing.',
                                                                           2, 5);

-- Уроки главы "DOM Manipulation" (chapter_id = 6)
INSERT INTO lessons (name, description, content, "order", chapter_id) VALUES
                                                                          ('Selecting and Modifying DOM Elements', 'Reading and changing page content using JavaScript.',
                                                                           'The Document Object Model, or DOM, represents an HTML page as a tree of nodes that JavaScript can read and modify. Elements can be selected using methods such as getElementById, querySelector, and querySelectorAll, which return one or more references to matching elements in the page. Once an element is selected, its content can be changed by modifying properties such as textContent or innerHTML, and its appearance can be adjusted by changing its class list or individual style properties. Because DOM manipulation directly affects what the user sees, excessive or repeated changes can impact page performance, which is why batching updates is often recommended.',
                                                                           1, 6),
                                                                          ('Handling Events in JavaScript', 'Responding to user interactions such as clicks and keyboard input.',
                                                                           'Event handling allows JavaScript to respond to user actions such as clicking a button, submitting a form, or pressing a key. The addEventListener method attaches a function to an element that runs whenever a specified event occurs, and multiple listeners can be attached to the same element without overwriting one another. The event object passed to a handler function contains details about the interaction, such as which key was pressed or which element triggered the event. Event propagation determines the order in which nested elements receive an event, moving from the outermost ancestor down to the target element during the capturing phase, and back up during the bubbling phase.',
                                                                           2, 6);

-- +goose Down

DELETE FROM lessons WHERE chapter_id IN (
    SELECT id FROM chapters WHERE course_id IN (
        SELECT id FROM courses WHERE name IN ('Golang Developer', 'Python Developer', 'JavaScript Developer')
    )
);

DELETE FROM chapters WHERE course_id IN (
    SELECT id FROM courses WHERE name IN ('Golang Developer', 'Python Developer', 'JavaScript Developer')
);

DELETE FROM courses WHERE name IN ('Golang Developer', 'Python Developer', 'JavaScript Developer');