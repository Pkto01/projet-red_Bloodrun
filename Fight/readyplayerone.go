package fight

def goblinPattern(goblin_name, target_name, max_hp):
    print(f"{goblin_name} commence le combat contre {target_name}.")
    for turn in range(1, 4):  # 3 tours de combat
        if turn % 2 == 1:  # Tour 1 et 3
            damage = 100
        else:  # Tour 2
            damage = 200
        current_hp = max_hp - (damage * turn)
        if current_hp < 0:
            current_hp = 0
        print(f"Tour {turn}: {goblin_name} inflige {damage} de dégâts à {target_name}. Points de vie restants : {current_hp}/{max_hp}")

def characterTurn():
    print("\n--- Tour du joueur ---")
    print("Menu :")
    print("1. Attaquer")
    print("2. Inventaire")
    choice = input("Choisissez une action (1 ou 2) : ")

    if choice == "1":
        attack_name = "Attaque basique"
        damage = 5
        target_max_hp = 50
        target_current_hp = target_max_hp - damage
        print(f"Vous utilisez '{attack_name}', infligez {damage} dégâts. Points de vie restants de l'adversaire : {target_current_hp}/{target_max_hp}")
    elif choice == "2":
        print("Vous consultez votre inventaire.")
    else:
        print("Choix invalide.")

# Exécution des fonctions
goblinPattern("Gobelin d'entraînement", "Personnage 5", 100)
characterTurn()