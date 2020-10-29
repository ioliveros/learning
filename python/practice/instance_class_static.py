"""Practicing concepts: 

    difference between

        instance method
        class method
        static method

"""

class Person():

    name = None
    age = None
    gender = None

    def __init__(self, name, age, gender):

        self.name = name
        self.age = age
        self.gender = gender

    """
        Class methods are like methods used as an alternative constructor in which you can change initial values without
        instantiating the class itself

        eg. Person.set_name('ian')

    """
    @classmethod
    def set_name(cls, name):
        age = None
        gender = None
        return cls(name=name, age=age, gender=gender)

    """
        Static methods are like contained methods within the class which don't take any depedencies
        from class (cls) and instance (self)        
    """
    @staticmethod
    def is_legal_age(age):        
        return True if age >= 21 else False

    def add_user_information(self, age):
        """adding user information then using a static method that checks if age is legal or not"""
        
        self.age = age

        if not self.is_legal_age(age):
            return f"{self.name} a {self.gender} is below legal age"
        
        return "OK"