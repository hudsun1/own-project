class CreatePayments < ActiveRecord::Migration[8.1]
  def change
    create_table :payments do |t|
      t.decimal :amount
      t.string :description

      t.timestamps
    end
  end
end
