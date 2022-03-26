package pers.fang.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.List;

/**
 * @author FANG
 * @create 2022-03-26 16:32
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Order implements Serializable {
    private Integer id;
    private String orderNo;
    private String orderAddress;
    private Double totalPrice;
    private List<Product> productList;
}
