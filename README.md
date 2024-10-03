# stori
![Workflow Status](https://github.com/gastonharari/stori/actions/workflows/main.yml/badge.svg)

To run the project, run this following commands:
```bash
docker-compose up --build
docker-compose run --rm processtransactions --file=/data/tsx.csv --email="email_to_send_summary_to@stori.com"
```

A better design that I would prefer to follow is more scalable, adheres more closely to the Single Responsibility Principle, and follows an event-driven architecture
![Screenshot 2024-10-03 195354](https://github.com/user-attachments/assets/bb2b4ade-08c2-44b9-b1c5-1b591e451ce0)

Another csv file that can be used to test the application is the `tsx2.csv` file in root directory.
