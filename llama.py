import requests


def main():
    session = requests.Session()
    query = input("\nPlease type the inputs for the model, then press Enter:\n\n")
    res = session.get(f"http://localhost:8080?text={query}", stream=True)
    try:
        print("\n\n--- streaming model output ---")
        for line in res.iter_content(3):
            try:
                x = line.decode()
                if x:
                    print(x, end="", flush=True)
            except KeyboardInterrupt:
                break
            except:
                continue
    except:
        ...
    finally:
        print("\n\n--- done ---")


if __name__ == "__main__":
    main()
