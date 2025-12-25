class CreateQrs < ActiveRecord::Migration[8.1]
  def change
    create_table :qrs do |t|
      t.string :imageUrl
      t.string :description

      t.timestamps
    end
  end
end
