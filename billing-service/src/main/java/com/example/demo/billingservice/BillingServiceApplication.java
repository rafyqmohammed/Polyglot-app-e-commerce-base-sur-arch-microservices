package com.example.demo.billingservice;

import com.example.demo.billingservice.entities.Bill;
import com.example.demo.billingservice.entities.ProductItem;
import com.example.demo.billingservice.repository.BillRepository;
import com.example.demo.billingservice.repository.ProductItemRepository;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.openfeign.EnableFeignClients;
import org.springframework.context.annotation.Bean;

import java.util.Date;

@SpringBootApplication
@EnableFeignClients
public class BillingServiceApplication {

    public static void main(String[] args) {
        SpringApplication.run(BillingServiceApplication.class, args);
    }

//    @Bean
//    CommandLineRunner init(ProductItemRepository productItemRepository, BillRepository billRepository) {
//        return args -> {
//            // insertion de bills
//            Bill bill1 = new Bill();
//            bill1.setBillingDate(new Date());
//            bill1.setCustomerId(1);
//            bill1 = billRepository.save(bill1);
//            Bill bill2 = new Bill();
//            bill2.setBillingDate(new Date());
//            bill2.setCustomerId(2);
//            bill2 = billRepository.save(bill2);
//
//            // insertion de productItems
//            // 1 bill1 => 2 productItems
//            ProductItem productItem1 = new ProductItem();
//            productItem1.setBill(bill1);
//            productItem1.setProductId(1);
//            productItem1.setQuantity(2);
//            productItem1.setUnitPrice(100);
//            productItemRepository.save(productItem1);
//            ProductItem productItem2 = new ProductItem();
//            productItem2.setBill(bill1);
//            productItem2.setProductId(2);
//            productItem2.setQuantity(1);
//            productItem2.setUnitPrice(200);
//            productItemRepository.save(productItem2);
//            // 1 bill2 => 1 productItem
//            ProductItem productItem3 = new ProductItem();
//            productItem3.setBill(bill2);
//            productItem3.setProductId(3);
//            productItem3.setQuantity(4);
//            productItem3.setUnitPrice(50);
//            productItemRepository.save(productItem3);
//
//        };
//    }
}
