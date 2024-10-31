/**
 * 类型转换工具类
 */
const Convert = {
    /**
     * 转换为整数
     * @param {*} value - 要转换的值
     * @param {number} defaultValue - 转换失败时的默认值
     * @returns {number}
     */
    toInt(value, defaultValue = 0) {
        if (value === null || value === undefined || value === '') {
            return defaultValue;
        }
        const num = parseInt(value, 10);
        return isNaN(num) ? defaultValue : num;
    },

    /**
     * 转换为浮点数
     * @param {*} value - 要转换的值
     * @param {number} defaultValue - 转换失败时的默认值
     * @param {number} decimals - 保留小数位数
     * @returns {number}
     */
    toFloat(value, defaultValue = 0, decimals = 2) {
        if (value === null || value === undefined || value === '') {
            return defaultValue;
        }
        const num = parseFloat(value);
        if (isNaN(num)) return defaultValue;
        return Number(num.toFixed(decimals));
    },

    /**
     * 转换为整数数组
     * @param {*} value - 要转换的值（可以是数组、逗号分隔的字符串等）
     * @param {number} defaultValue - 转换失败时的默认值
     * @returns {number[]}
     */
    toIntArray(value, defaultValue = 0) {
        if (!value) return [];

        // 如果已经是数组
        if (Array.isArray(value)) {
            return value.map(item => this.toInt(item, defaultValue));
        }

        // 如果是字符串，按逗号分割
        if (typeof value === 'string') {
            return value.split(',')
                .map(item => item.trim())
                .filter(item => item !== '')
                .map(item => this.toInt(item, defaultValue));
        }

        return [this.toInt(value, defaultValue)];
    },

    /**
     * 转换表单数据中的指定字段为整数
     * @param {Object} formData - 表单数据对象
     * @param {string[]} intFields - 需要转换为整数的字段名数组
     * @param {number} defaultValue - 转换失败时的默认值
     * @returns {Object} 转换后的新对象
     */
    convertFormFields(formData, intFields = [], defaultValue = 0) {
        const result = { ...formData };
        intFields.forEach(field => {
            if (field in result) {
                result[field] = this.toInt(result[field], defaultValue);
            }
        });
        return result;
    },

    /**
     * 转换表单数据中的数组字段
     * @param {Object} formData - 表单数据对象
     * @param {string[]} arrayFields - 需要转换为数组的字段名数组
     * @param {number} defaultValue - 转换失败时的默认值
     * @returns {Object} 转换后的新对象
     */
    convertArrayFields(formData, arrayFields = [], defaultValue = 0) {
        const result = { ...formData };
        arrayFields.forEach(field => {
            if (field in result) {
                result[field] = this.toIntArray(result[field], defaultValue);
            }
        });
        return result;
    }
};