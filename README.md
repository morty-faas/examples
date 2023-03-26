# Morty function examples

This repository is a monorepo that contains function examples that can be deployed in seconds on a Morty instance.

Each folder of this repository is a function.

To build it for your Morty instance, ensure you've configured your context, and execute the following command :

```bash
morty fn build $function
```

Where `$function` is one of the folder here.

You can now invoke your function using the following command :

```bash
morty fn invoke $function
```
