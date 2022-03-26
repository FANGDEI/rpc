package pers.fang.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;

/**
 * @author FANG
 * @create 2022-03-26 16:07
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class Product implements Serializable {
    private Integer id;
    private String productName;
    private Integer productNum;
    private Double productPrice;
}
