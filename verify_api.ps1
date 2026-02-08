$baseUrl = "http://localhost:8080/api"

Write-Host "--- Testing Search By Name ---"
Invoke-RestMethod -Uri "$baseUrl/produk?name=Kopi" -Method Get | ConvertTo-Json -Depth 5

Write-Host "`n--- Testing Checkout ---"
$checkoutPayload = @{
    items = @(
        @{ product_id = 1; quantity = 2 }
    )
}
$transaction = Invoke-RestMethod -Uri "$baseUrl/checkout" -Method Post -Body ($checkoutPayload | ConvertTo-Json) -ContentType "application/json"
$transaction | ConvertTo-Json -Depth 5

Write-Host "`n--- Testing Today's Report ---"
Invoke-RestMethod -Uri "$baseUrl/report/hari-ini" -Method Get | ConvertTo-Json -Depth 5

Write-Host "`n--- Testing Date Range Report ---"
Invoke-RestMethod -Uri "$baseUrl/report?start_date=2026-01-01&end_date=2026-12-31" -Method Get | ConvertTo-Json -Depth 5
