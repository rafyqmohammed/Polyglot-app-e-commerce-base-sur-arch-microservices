package com.example.demo.billingservice.feign;


import com.example.demo.billingservice.model.Customer;
import com.example.demo.billingservice.model.Product;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.hateoas.PagedModel;


@FeignClient(name = "inventory-service")
public interface ProductRestClient {
    @GetMapping("products/{id}")
    Product getProductById(@PathVariable int id);
    @GetMapping("/products")
    PagedModel<Product> getAllProducts();
}