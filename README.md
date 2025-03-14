# Sales-Insights

## Running the Application

This application is built with **Go 1.24.1** (tested on `linux/amd64`) and uses **SQLite 3.46.0** as the database. Follow the steps below to run it after cloning the repository.

### Prerequisites
- **Go**: Version 1.24.1 or later (`go version` to check).
- **SQLite**: Version 3.46.0 or later (`sqlite3 --version` to check).

### Steps to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/singhamritpalAP/Sales-Insights.git
   
2. Navigate to the cmd directory:
   ```bash
   cd cmd
3. Run the application:
   ```bash
   go run .
The application will start a server on http://localhost:8080


## API Endpoints

| Route                                                            | Method | Body | Sample Response                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Description                                                                                                  |
|------------------------------------------------------------------|--------|------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------|
| `/refresh`                                                       | POST   | None | ```"Data refreshed successfully."```                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Triggers a refresh of the database by loading and processing data from the CSV file.                         |
| `/top-products/overall?n={n}&start_date={start}&end_date={end}`  | GET    | None | ```      [{"ProductID":"P123","ProductName":"UltraBoost Running Shoes","Category":"Shoes","UnitPrice":180},{"ProductID":"P456","ProductName":"iPhone 15 Pro","Category":"Electronics","UnitPrice":1299},{"ProductID":"P789","ProductName":"Levi's 501 Jeans","Category":"Clothing","UnitPrice":59.99}]```                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Retrieves the top `n` products by total quantity sold across all categories within the specified date range. |
| `/top-products/category?n={n}&start_date={start}&end_date={end}` | GET    | None | ``` {"Clothing":[{"ProductID":"P789","ProductName":"Levi's 501 Jeans","Category":"Clothing","UnitPrice":59.99}],"Electronics":[{"ProductID":"P456","ProductName":"iPhone 15 Pro","Catego ry":"Electronics","UnitPrice":1299},{"ProductID":"P234","ProductName":"Sony WH-1000XM5 Headphones","Category":"Electronics","UnitPrice":349.99}],"Shoes":[{"ProductID":"P123","ProductName":"UltraBoost Running Shoes","Category":"Shoes","UnitPrice":180}]} ```                                                                                                                                                                                                                                                                                                                                      | Retrieves the top `n` products per category by quantity sold within the specified date range.                |
| `/top-products/region?n={n}&start_date={start}&end_date={end}`   | GET    | None | ``` {"Asia":[{"ProductID":"P789","ProductName":"Levi's 501 Jeans","Category":"Clothing","UnitPrice":59.99},{"ProductID":"P456","ProductName":"iPhone 15 Pro","Category":"Electronics","U nitPrice":1299}],"Europe":[{"ProductID":"P456","ProductName":"iPhone 15 Pro","Category":"Electronics","UnitPrice":1299}],"North America":[{"ProductID":"P123","ProductName":"UltraBo ost Running Shoes","Category":"Shoes","UnitPrice":180},{"ProductID":"P234","ProductName":"Sony WH-1000XM5 Headphones","Category":"Electronics","UnitPrice":349.99}],"South America":[{"ProductID":"P123","ProductName":"UltraBoost Running Shoes","Category":"Shoes","UnitPrice":180}]} ``` | Retrieves the top `n` products per region by quantity sold within the specified date range.                  |

### Usage Examples

#### Refresh Database
```bash
curl -X POST http://localhost:8080/refresh
```
#### Get Top Products Overall
```bash
curl "http://localhost:8080/top-products/overall?n=3&start_date=2023-01-01&end_date=2024-12-31"
```
#### Get Top Products by Category
```bash
curl "http://localhost:8080/top-products/category?n=2&start_date=2024-01-01&end_date=2024-06-30"
```
#### Get Top Products by Region
```bash
curl "http://localhost:8080/top-products/region?n=5&start_date=2023-01-01&end_date=2024-12-31"
```
